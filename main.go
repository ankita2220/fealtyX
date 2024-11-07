// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/gorilla/mux"
// )

// // Define the Student struct
// type Student struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name" validate:"required"`
// 	Age   int    `json:"age" validate:"gte=18,lte=100"`
// 	Email string `json:"email" validate:"required,email"`
// }

// // In-memory data store for students
// var (
// 	students  = make(map[int]Student)
// 	idCounter = 1
// 	validate  = validator.New()
// )

// // Main function to initialize the server and routes
// func main() {
// 	// Create a new router using Gorilla Mux
// 	router := mux.NewRouter()

// 	// Define the routes for CRUD operations
// 	router.HandleFunc("/students", CreateStudent).Methods("POST")
// 	router.HandleFunc("/students", GetAllStudents).Methods("GET")
// 	router.HandleFunc("/students/{id}", GetStudentByID).Methods("GET")
// 	router.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
// 	router.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")

// 	// Start the HTTP server on port 3031
// 	fmt.Println("Server running at http://localhost:3031")
// 	log.Fatal(http.ListenAndServe(":3031", router))
// }

// // Create a new student
// func CreateStudent(w http.ResponseWriter, r *http.Request) {
// 	var student Student
// 	err := json.NewDecoder(r.Body).Decode(&student)
// 	if err != nil || validate.Struct(student) != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	// Assign ID and store in the in-memory map
// 	student.ID = idCounter
// 	students[idCounter] = student
// 	idCounter++

// 	// Respond with the created student
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(student)
// }

// // Get all students
// func GetAllStudents(w http.ResponseWriter, r *http.Request) {
// 	studentList := make([]Student, 0, len(students))
// 	for _, student := range students {
// 		studentList = append(studentList, student)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(studentList)
// }

// // Get a student by ID
// func GetStudentByID(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	idStr := params["id"]
// 	id, err := strconv.Atoi(idStr) // Convert string to int
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	student, exists := students[id]
// 	if !exists {
// 		http.Error(w, "Student not found", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(student)
// }

// // Update a student by ID
// func UpdateStudent(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	idStr := params["id"]
// 	id, err := strconv.Atoi(idStr) // Convert string to int
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	var student Student
// 	err = json.NewDecoder(r.Body).Decode(&student)
// 	if err != nil || validate.Struct(student) != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	// Update student information
// 	existingStudent, exists := students[id]
// 	if !exists {
// 		http.Error(w, "Student not found", http.StatusNotFound)
// 		return
// 	}

// 	existingStudent.Name = student.Name
// 	existingStudent.Age = student.Age
// 	existingStudent.Email = student.Email
// 	students[id] = existingStudent

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(existingStudent)
// }

// // Delete a student by ID
// func DeleteStudent(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	idStr := params["id"]
// 	id, err := strconv.Atoi(idStr) // Convert string to int
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	_, exists := students[id]
// 	if !exists {
// 		http.Error(w, "Student not found", http.StatusNotFound)
// 		return
// 	}

// 	delete(students, id)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
// }

// // Get a summary for a student by ID
// func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	idStr := params["id"]
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	student, exists := students[id]
// 	if !exists {
// 		http.Error(w, "Student not found", http.StatusNotFound)
// 		return
// 	}

// 	// Call the Ollama API to generate a summary
// 	summary, err := GenerateOllamaSummary(student)
// 	if err != nil {
// 		http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{"summary": summary})
// }

// // Function to call Ollama API and generate a summary for a student
// func GenerateOllamaSummary(student Student) (string, error) {
// 	// Define the payload for the Ollama request
// 	payload := map[string]string{
// 		"prompt": fmt.Sprintf("Generate a profile summary for a student named %s, age %d, email %s.", student.Name, student.Age, student.Email),
// 	}

// 	// Marshal the payload to JSON
// 	jsonData, err := json.Marshal(payload)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Make the HTTP request to Ollama API
// 	resp, err := http.Post("http://localhost:1140/generate", "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	// Read the response from Ollama API
// 	var result map[string]interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return "", err
// 	}

// 	// Extract the summary text
// 	summary, ok := result["summary"].(string)
// 	if !ok {
// 		return "", fmt.Errorf("invalid response from Ollama API")
// 	}

// 	return summary, nil
// }

package main

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

// Define the Student struct
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"gte=18,lte=100"`
	Email string `json:"email" validate:"required,email"`
}

// In-memory data store for students
var (
	students  = make(map[int]Student)
	idCounter = 1
)

// Main function to initialize the server and routes
func main() {
	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Define the routes for CRUD operations
	router.HandleFunc("/students", CreateStudent).Methods("POST")
	router.HandleFunc("/students", GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", GetStudentByID).Methods("GET")
	router.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")
	router.HandleFunc("/students/{id}/summary", GetStudentSummary).Methods("GET") // New route for summary

	// Start the HTTP server on port 3031
	fmt.Println("Server running at http://localhost:3031")
	log.Fatal(http.ListenAndServe(":3031", router))
}

// Create a new student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Assign ID and store in the in-memory map
	student.ID = idCounter
	students[idCounter] = student
	idCounter++

	// Respond with the created student
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// Get all students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	studentList := make([]Student, 0, len(students))
	for _, student := range students {
		studentList = append(studentList, student)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentList)
}

// Get a student by ID
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr) // Convert string to int
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	student, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// Update a student by ID
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr) // Convert string to int
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var student Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update student information
	existingStudent, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	existingStudent.Name = student.Name
	existingStudent.Age = student.Age
	existingStudent.Email = student.Email
	students[id] = existingStudent

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingStudent)
}

// Delete a student by ID
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr) // Convert string to int
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	_, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	// Delete student
	delete(students, id)
	w.WriteHeader(http.StatusNoContent)
}

// Get summary for a specific student
func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	student, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	// Generate summary using Ollama API
	summary, err := GenerateOllamaSummary(student)
	if err != nil {
		http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
		return
	}

	// Clean the summary output to remove ANSI sequences
	cleanSummary := removeANSISequences(summary)

	// Return summary in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"summary": cleanSummary})
}

// Generate student summary using Ollama API
func GenerateOllamaSummary(student Student) (string, error) {
	os.Setenv("HOME", "/Users/mansibakshi") // Adjust path as needed
	prompt := fmt.Sprintf("Generate a detailed profile summary for a student named %s, age %d, and email %s.", student.Name, student.Age, student.Email)
	cmd := exec.Command("ollama", "run", "llama3:latest", prompt)
	cmd.Env = append(os.Environ())

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v", err)
	}
	return string(output), nil
}

// Remove ANSI escape sequences from the summary
func removeANSISequences(input string) string {
	re := regexp.MustCompile(`\x1B\[[0-?]*[ -/]*[@-~]`)
	input = re.ReplaceAllString(input, "")
	input = strings.ReplaceAll(input, "\n", " ")
	input = strings.ReplaceAll(input, "\r", "")
	return input
}
