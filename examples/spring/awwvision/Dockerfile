FROM openjdk:8

EXPOSE 8080

WORKDIR /awwvision
COPY target/*.jar .

ENTRYPOINT ["java", "-jar", "awwvision-spring-0.0.1-SNAPSHOT.jar"]
