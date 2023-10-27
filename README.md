## How to Use

1. **Clone the Repository**: Start by cloning the project repository from Git:

    ```shell
    git clone git@github.com:Little-BlackCat/check-password.git
    ```

2. **Run Docker Compose**: Navigate to the project directory and run Docker Compose to start the services (nginx, Golang service, and PostgreSQL):

    ```shell
    docker-compose up -d
    ```

3. **Submit a POST Request**: Use `curl` to submit a POST request to the Golang service with your desired initial password. Replace `<password>` with your password:

    ```shell
    curl -X POST -H "Content-Type: application/json" -d '{"init_password": "<password>"}' http://localhost:8080/api/check_password
    ```

4. **Check the Database**: If you want to check the PostgreSQL database, you can access it using the following command:

    ```shell
    docker exec -it check-password-db-1 psql -U postgres -d postgres
    ```

    Once you're inside the PostgreSQL shell, you can retrieve data from the `password_log` table:

    ```sql
    SELECT * FROM password_log;
    ```

5. **Run Tests**: To run the tests for the Golang service, use the following command:

    ```shell
    go test
    ```
