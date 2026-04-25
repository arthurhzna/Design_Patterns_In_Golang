# Design Patterns in Golang 🧠🐹

summary and quick notes for design patterns practiced in this repository  🚀

## 🧭 Pattern Categories at a Glance

### Creational (Object Creation)
- Focuses on how objects are created.
- Goal: make object instantiation flexible, safe, and less error-prone.
- Typical question: "How should this object be created?"

### Structural (Object Composition)
- Focuses on how objects and types are composed into larger structures.
- Goal: keep relationships clean, compatible, and maintainable.
- Typical question: "How should these objects be connected or wrapped?"

### Behavioral (Object Interaction)
- Focuses on how objects communicate and share responsibilities.
- Goal: organize flow, algorithms, state changes, and collaboration.
- Typical question: "Who does what, when, and how?"

## 🏗️ Creational Patterns

### Builder 🧱
- Separate component for when object construction gets too complicated.
- Can create mutually cooperating sub-builders.
- Often has a fluent interface.

### Prototype 🧬
- Creation of an object from an existing object.
- Requires either explicit deep copy or copy through serialization.

### Factory 🏭
- Factory functions (constructors) are common.
- A factory can be a simple function or a dedicated struct.

### Singleton 1️⃣
- Use when you need to ensure only a single instance exists.
- Can be made thread-safe and lazy.
- Consider extracting an interface or using dependency injection.

## 🧩 Structural Patterns

### Adapter 🔌
- Converts the interface you get to the interface you need.

### Bridge 🌉
- Decouples abstraction from implementation.

### Composite 🌲
- Allows clients to treat individual objects and compositions uniformly.

### Decorator 🎁
- Attaches additional responsibilities to objects.
- Can be done through embedding or pointers.

### Facade 🧰
- Provides a single unified interface over a set of interfaces.

### Flyweight 🍃
- Efficiently supports very large numbers of similar objects.

### Proxy 🛡️
- Provides a surrogate object that forwards calls to the real object while performing additional functions.
- Example: access control, communication, logging, etc.

## 🎭 Behavioral Patterns

### Chain of Responsibility ⛓️
- Allows components to process information/events in a chain.
- Each element in the chain refers to the next element, or you can keep a list and iterate through it.

### Command 🕹️
- Encapsulates a request into a separate object.
- Good for audit, replay, and undo/redo.
- Part of CQS/CQRS style thinking.

### Interpreter 📖
- Transforms textual input into structures (e.g. ASTs).
- Used by interpreters, compilers, static analysis tools, etc.
- Compiler theory is a separate branch of computer science.

### Iterator 🔁
- Provides an interface for accessing elements of an aggregate object.

### Mediator 🤝
- Provides mediation services between several objects.
- Example: message passing, chat room.

### Memento 🧠💾
- Yields tokens representing system states.
- Tokens do not allow direct manipulation, but can be used in appropriate APIs.

### Observer 👀
- Allows notifications of changes/happenings in a component.

### State 🔄
- Models systems by having one of a set of possible states and transitions between them.
- Such a system is called a state machine.
- Special frameworks exist to orchestrate state machines.

### Strategy and Template Method ♟️🧩
- Both define a skeleton algorithm with details filled in by the implementer.
- Strategy uses composition; Template Method does not.

### Visitor 🚶
- Allows non-intrusive addition of functionality to hierarchies.

