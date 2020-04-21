# Temperature control system

This is a simple project to demostrate Go-Micro framework. This repository consists of three diffrent project that simulate IoT system.

The projects are:
    -TemperatureControlSystem
    -AirCondtioner
    -Thermometar


To compile all three projects simply run ` go build ` command inside all three folders of project and than run .exe file. You can run multiple instances of every project expect for TemepratureControlSystem (because of the prot assignemnt).

## TemperatureControlSystem

This is a web controller that runs on localhost:8080 address. You can run only one instance per computer. It gives you a  "Dashboard" where you can see all the rooms, their temperature and status of each air condtioner. You can also set your own status of aircondtioner.

## AirCondtioner

AirCondtioner changes the room temperature, or rather the temperature stored in thermometar, based on settings. If the power is off it doesn't do anything. If AutoMode is on it trys to keep the temperature closes possible to desired temperature. Othewerise it heats/cools the room for random number bettween 0 and 5 multiply by current power of AirCondtioner.

## Thermometar

Thermometar is in charge of tracking the temperature of the room. Like AirCondtioner you can only run one instance of the project per room. It also changes the room temperature randomly so all features of AirCondtioner can be demostrated.