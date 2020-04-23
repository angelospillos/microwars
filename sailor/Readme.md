# Naughty Sailor
Java Micronaut with GraalVM native image.

## GraalVM Native Image
First get a fat JAR going:
```
./gradlew assemble
```

Assuming GraalVM _and_ native-image are installed:
```
native-image --no-server -cp build/libs/sailor-*-all.jar
```