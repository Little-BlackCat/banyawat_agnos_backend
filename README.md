## How to Use

1. Clone the repository first:

    ```shell
    git clone git@github.com:Little-BlackCat/check-password.git
    ```

2. Start the Docker Compose environment containing Nginx, Golang service, and PostgreSQL:

    ```shell
    docker-compose up -d
    ```

3. Use `curl` to send a POST request to the service with the desired password for initialization:

    ```shell
    curl -X POST -H "Content-Type: application/json" -d '{"init_password": <password>}' http://localhost:8080/api/check_password
    ```

4. To check the database, you can access the PostgreSQL container using the following command:

    ```shell
    docker exec -it strong-password-service-db-1 psql -U postgres -d postgres
    ```

5. Once connected to the PostgreSQL database, you can view the saved records in the `password_log` table:

    ```sql
    select * from password_log;
    ```