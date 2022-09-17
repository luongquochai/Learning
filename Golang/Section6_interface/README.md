THIS IS A INTERFACE NOTE

- Interfaces are not generic types: Other language have 'generic' types - go does not.
- Interfaces are 'implicit': we dont manually have to say that our custom type satisfies some interface.
- Interfaces are a contract to help us manage types: GARBAGE IN -> GARBAGE OUT. If our custome type's implementation of a function is broken then interfaces wont help us !
- Interfaces are tough (hard). Step #1 is understanding how to read them: Understand how to read interfaces in the standard library. Writing your own interfaces is tough and requires experience.