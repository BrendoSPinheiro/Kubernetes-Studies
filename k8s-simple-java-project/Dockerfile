FROM maven:3.9.6-eclipse-temurin-21 AS build

WORKDIR /app

COPY ./pom.xml .
COPY ./src src

RUN mvn dependency:go-offline -B
RUN mvn clean package -DskipTests

FROM eclipse-temurin:21-jre

WORKDIR /app

COPY --from=build /app/target/*.jar app.jar

EXPOSE 8080

ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -jar app.jar"]