GRASP (General Responsibility Assignment Software Principles) patterns are a set of principles aimed at guiding the assignment of responsibilities to classes and objects in object-oriented design. These patterns help developers create flexible, maintainable, and cohesive software systems. Below is an overview of the key GRASP patterns:

### 1. [Information Expert](https://github.com/NikolaiKovalenko/edu-grasp-patterns/tree/main/Information_Expert)
- **Definition**: Assign responsibility to the class that has the necessary information to fulfill it.
- **Purpose**: Promotes low coupling and high cohesion by ensuring that methods are allocated to classes that have the necessary data.
- **Example**: A `Customer` class should handle tasks related to customer data, such as validating a customer’s credit because it owns that data.

### 2. [Creator](https://github.com/NikolaiKovalenko/edu-grasp-patterns/tree/main/Creator)
- **Definition**: Assign the responsibility of creating an instance of a class to a class that aggregates, contains, or closely uses the instance.
- **Purpose**: Helps reduce dependencies among classes and fosters more intuitive designs by ensuring that classes responsible for an object also create it.
- **Example**: A `ShoppingCart` class should create `CartItem` instances, as it contains and manages cart items.

### 3. [Controller](https://github.com/NikolaiKovalenko/edu-grasp-patterns/tree/main/Controller)
- **Definition**: Assign the responsibility of handling system events to a controller class that represents the overall system or a use case scenario.
- **Purpose**: Delegates responsibility from user interface components to a dedicated controller, facilitating the separation of concerns.
- **Example**: In a GUI application, the `MainController` might handle events like user input, delegating functionality to model and view components.

### 4. Low Coupling
- **Definition**: Aim to reduce dependencies between classes, promoting loose coupling.
- **Purpose**: Ensures that changes in one class minimally affect others, enhancing maintainability and flexibility.
- **Example**: Using interfaces to allow a `PaymentProcessor` to switch between different payment methods without modifying the underlying logic.

### 5. High Cohesion
- **Definition**: Assign related responsibilities to a single class to enhance the cohesion of that class.
- **Purpose**: Makes classes easier to understand and maintain by focusing on a single set of related tasks.
- **Example**: A `ReportGenerator` class should exclusively handle tasks related to report generation, avoiding unrelated functionalities.

### 6. Polymorphism
- **Definition**: Use polymorphic methods to handle variation in behavior, often via interfaces or abstract classes.
- **Purpose**: Reduces the need for conditionals in the code and simplifies class design by leveraging the benefits of polymorphism.
- **Example**: A `Shape` interface implemented by `Circle` and `Square` classes may have a method `draw()`, allowing for polymorphic invocation for different shapes.

### 7. Pure Fabrication
- **Definition**: Create a class that does not represent a concept in the problem domain to achieve low coupling, high cohesion, or reuse.
- **Purpose**: Helps maintain a clean design by avoiding an unnecessary dependence on a real-world concept, often leading to a cleaner architecture.
- **Example**: A `NotificationService` class might handle email and SMS notifications without being tied directly to the domain model.

### 8. Indirection
- **Definition**: Assign responsibilities in a way that avoids direct coupling between classes by introducing an intermediate class.
- **Purpose**: Reduces the dependencies between classes and aids in decoupling.
- **Example**: Using a service layer as an intermediary between data access and business logic components.

### 9. Protected Variance
- **Definition**: Allow subclasses to alter behavior while keeping their interfaces stable by using encapsulation.
- **Purpose**: Enables safe variations in behavior through inheritance while maintaining the client interface.
- **Example**: An abstract `Dialog` class with a method `createButton()` can be overridden in subclasses to produce various types of buttons.

### Summary

The GRASP principles provide a foundation for designing robust, maintainable, and scalable object-oriented systems. They encourage assigning responsibilities thoughtfully, balancing between cohesion and coupling, and utilizing polymorphism effectively. Applying these principles can lead to high-quality software designs that are easier to manage and evolve over time.