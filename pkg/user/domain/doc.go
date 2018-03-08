/*
Package domain is the heart layer of the software, and this is where the interesting stuff happens.
There is one package per aggregate and to each aggregate belongs entities, value objects, domain events, a repository interface and sometimes factories.
The core of the business logic, such as determining whether a handling event should be registered.
The structure and naming of aggregates, classes and methods in the domain layer should follow the ubiquitous language,
and you should be able to explain to a domain expert how this part of the software works by drawing a few simple diagrams and using the actual class and method names of the source code.
*/
package domain
