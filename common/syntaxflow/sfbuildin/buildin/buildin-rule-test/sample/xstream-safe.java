package com.example.springboot.controller;

import com.example.springboot.model.Person;
import com.example.springboot.model.PhoneNumber;
import com.thoughtworks.xstream.XStream;
import com.thoughtworks.xstream.io.xml.StaxDriver;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PersonController {

    @GetMapping("/person")
    public String getPerson() {
        Person person = new Person("John", "Doe");
        person.setPhone(new PhoneNumber(123, "1234-567"));

        XStream xstream = new XStream(new StaxDriver());
        xstream.alias("person", Person.class);
        xstream.alias("phonenumber", PhoneNumber.class);

        return xstream.toXML(person);
    }

    @PostMapping("/person")
    public ResponseEntity<Person> createPerson(@RequestBody String xml) {
        XStream xstream = new XStream(new StaxDriver());
        xstream.alias("person", Person.class);
        xstream.alias("phonenumber", PhoneNumber.class);

        xstream.setMode(XStream.NO_REFERENCE);

        Person person = (Person) xstream.fromXML(xml);
        return new ResponseEntity<>(person, HttpStatus.CREATED);
    }

    @PostMapping("/person2")
    public ResponseEntity<Person> createPerson2(@RequestBody String xml) {
        XStream xstream = new XStream(new StaxDriver());
        xstream.alias("person", Person.class);
        xstream.alias("phonenumber", PhoneNumber.class);

        xstream.allowTypes(new Class[]{Person.class});
        Person person = (Person) xstream.fromXML(xml);
        return new ResponseEntity<>(person, HttpStatus.CREATED);
    }
}