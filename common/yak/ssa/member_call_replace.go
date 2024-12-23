package ssa

import (
	"github.com/yaklang/yaklang/common/log"
	"github.com/yaklang/yaklang/common/utils"
)

// ReplaceMemberCall replace all member or object relationship
// and will fixup method function call
func ReplaceMemberCall(v, to Value) map[string]Value {
	ret := make(map[string]Value)
	builder := v.GetFunc().builder
	recoverScope := builder.SetCurrent(v)
	defer recoverScope()
	createPhi := generatePhi(builder, nil, nil)

	// replace object member-call
	if v.IsObject() {
		for _, member := range v.GetAllMember() {
			// replace this member object to to
			key := member.GetKey()
			// remove this member from v
			v.DeleteMember(key)

			// create member of `to` value with key
			// if fun := GetMethod(value.GetType(), key.String()); fun != nil {
			// 	return NewClassMethod(fun, value)
			// }
			// re-set type
			res := checkCanMemberCallExist(to, key)
			name, typ := res.name, res.typ
			// toMember := builder.getOriginMember(name, typ, to, key)
			toMember := builder.PeekValue(name)

			// then, we will replace value, `member` to `toMember`
			if member.GetOpcode() != SSAOpcodeUndefined {
				member.SetName(name)
				member.SetType(typ)
				setMemberCallRelationship(to, key, member)
				log.Warn("ReplaceMemberCall can create phi, but we cannot find cfgEntryBlock")
				if utils.IsNil(toMember) {
					ret[name] = member
				} else {
					ret[name] = createPhi(name, []Value{toMember, member})
				}
			}
			if utils.IsNil(toMember) {
				toMember = builder.ReadMemberCallValue(to, key)
			}
			ReplaceAllValue(member, toMember)
			DeleteInst(member)

			ret[name] = toMember
		}
	}

	// TODO : this need more test, i think this code error
	if v.IsMember() {
		obj := v.GetObject()
		setMemberCallRelationship(obj, v.GetKey(), v)
	}
	return ret
}
