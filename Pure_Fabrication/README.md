### Pure Fabrication Pattern

The **Pure Fabrication** pattern is a GRASP (General Responsibility Assignment Software Principles) pattern that suggests creating a class that does not represent a concept in the problem domain to achieve low coupling, high cohesion, or reuse. A Pure Fabrication is an artificial construct made to support these principles, often introduced to handle behaviors that do not naturally belong to any of the domain objects.

#### Key Features:
- **Non-Domain Class**: A Pure Fabrication does not correspond to a real-world entity or domain concept.
- **Behavior Focus**: Provides behavior-focused services that are needed throughout the system.
- **Decoupling**: Helps decouple responsibilities in the system that don't align well with domain objects.

#### Benefits:
- **High Cohesion**: Concentrates on providing specific services, avoiding bloating domain objects with unrelated responsibilities.
- **Reusability**: Can be reused across different parts of the system or even in different systems because it provides general-purpose functionality.
- **Low Coupling**: Allows domain objects to focus on their core responsibilities by offloading unrelated actions to the Pure Fabrication.

#### Potential Problems:
- **Complexity**: Introducing too many fabricated classes can lead to complexity and potentially unnecessary abstraction.
- **Domain-Model Blur**: Overuse may blur the application’s domain model, making it less intuitive to understand.

### Summary:
The Pure Fabrication pattern is useful when specific functionalities need to be extracted from domain objects to avoid clutter and improve cohesion. By introducing a class dedicated to these responsibilities (such as logging), the system maintains high cohesion, low coupling, and enhanced reusability without distorting the domain model.