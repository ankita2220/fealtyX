
### Project Overview
This project is a simple RESTful API for managing student records, built in Go. It supports CRUD operations for student data and generates a detailed profile summary using an external API. 

---

### Key Components of the Code

1. **Packages and Imports**
   - Imports essential packages for:
     - JSON handling (`encoding/json`)
     - HTTP server management (`net/http`)
     - Logging (`log`)
     - Command execution (`os/exec`)
     - Regular expressions (`regexp`) and string handling (`strconv`, `strings`)
     - Routing using **Gorilla Mux** (`github.com/gorilla/mux`), which is necessary for defining clean URL endpoints for each route.

   ```go
   import (
       "encoding/json"
       "fmt"
       "log"
       "net/http"
       "os"
       "os/exec"
       "regexp"
       "strconv"
       "strings"
       "github.com/gorilla/mux"
   )
   ```

2. **Student Struct Definition**
   - The `Student` struct defines the schema for each student record with four fields: `ID`, `Name`, `Age`, and `Email`.
   - Validation tags (e.g., `required`, `gte`, `lte`) are added to ensure data integrity when validation libraries are integrated.

   ```go
   type Student struct {
       ID    int    `json:"id"`
       Name  string `json:"name" validate:"required"`
       Age   int    `json:"age" validate:"gte=18,lte=100"`
       Email string `json:"email" validate:"required,email"`
   }
   ```

3. **In-Memory Data Store**
   - The `students` map acts as a simple database to store students by their ID.
   - `idCounter` is a global variable that ensures each new student gets a unique ID.

   ```go
   var (
       students  = make(map[int]Student)
       idCounter = 1
   )
   ```

4. **Router Initialization and Route Definitions**
   - The `mux.NewRouter()` function creates a new router instance. The following routes are defined for handling CRUD operations and summary generation:
     - **POST** `/students`: Create a new student record.
     - **GET** `/students`: Retrieve all students.
     - **GET** `/students/{id}`: Retrieve a specific student by ID.
     - **PUT** `/students/{id}`: Update a student's record.
     - **DELETE** `/students/{id}`: Delete a student record.
     - **GET** `/students/{id}/summary`: Generate a profile summary.

   ```go
   func main() {
       router := mux.NewRouter()
       router.HandleFunc("/students", CreateStudent).Methods("POST")
       router.HandleFunc("/students", GetAllStudents).Methods("GET")
       router.HandleFunc("/students/{id}", GetStudentByID).Methods("GET")
       router.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
       router.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")
       router.HandleFunc("/students/{id}/summary", GetStudentSummary).Methods("GET")
       
       fmt.Println("Server running at http://localhost:3031")
       log.Fatal(http.ListenAndServe(":3031", router))
   }
   ```

5. **CRUD Functions**
   - Each CRUD function corresponds to a route:
     - **CreateStudent**: Accepts a JSON object for a new student, assigns a unique ID, and adds it to the `students` map.
     - **GetAllStudents**: Returns a list of all students in JSON format.
     - **GetStudentByID**: Fetches a single student by their ID from the `students` map.
     - **UpdateStudent**: Modifies an existing student record.
     - **DeleteStudent**: Removes a student from the `students` map.

   ```go
   func CreateStudent(w http.ResponseWriter, r *http.Request) { ... }
   func GetAllStudents(w http.ResponseWriter, r *http.Request) { ... }
   func GetStudentByID(w http.ResponseWriter, r *http.Request) { ... }
   func UpdateStudent(w http.ResponseWriter, r *http.Request) { ... }
   func DeleteStudent(w http.ResponseWriter, r *http.Request) { ... }
   ```

6. **Generating a Summary for a Student**
   - **GetStudentSummary**: Calls `GenerateOllamaSummary` to create a profile summary for a given student, making use of the `ollama` command-line tool.
   - **GenerateOllamaSummary**: Uses the `exec.Command` function to run an external shell command, generating the summary.
   - **removeANSISequences**: Strips ANSI escape codes from the output for clean JSON display.

   ```go
   func GetStudentSummary(w http.ResponseWriter, r *http.Request) { ... }
   func GenerateOllamaSummary(student Student) (string, error) { ... }
   func removeANSISequences(input string) string { ... }
   ```

---

### Running the Application
1. Clone the repository or copy the files into your project directory.
2. Install dependencies, particularly **Gorilla Mux**.
3. Run the application:

   ```bash
   go run main.go
   ```

4. Access the server at `http://localhost:3031` and use an API client like Postman to test each endpoint.
5. Images
   
<img width="1440" alt="Screenshot 2024-11-08 at 2 45 28 AM" src="https://github.com/user-attachments/assets/9b7f2e58-999a-4fd9-8c87-23c5cb26db3f">
<img width="1440" alt="Screenshot 2024-11-08 at 2 44 36 AM" src="https://github.com/user-attachments/assets/38bdff9b-8e23-4cda-89fc-69defdf0fdce">
<img width="1440" alt="Screenshot 2024-11-08 at 2 45 25 AM" src="https://github.com/user-attachments/assets/8ec34245-403c-47d1-a876-194c4849786d">
--
