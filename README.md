    The cmd/server/ folder contains the main entry point for your application.
    The internal/ folder contains your application's core logic, divided into several packages.
    The app/ package contains the central home automation controller, mediator, and user-related functions.
    The devices/ package is further divided into sub-packages for each device type (light, speaker, thermostat, and security). Each sub-package contains the implementation of the specific device, along with the related design patterns (Factory Method, Abstract Factory, etc.).
    Optionally, you can create a separate patterns/ package to keep the implementations of design patterns that are used across multiple devices.
    The api/ package contains API-related functions and handlers, which can be further divided into sub-packages based on the API type (REST, gRPC, etc.).
    The pkg/ folder is used for external packages, such as utility functions and helpers that can be shared across multiple projects.