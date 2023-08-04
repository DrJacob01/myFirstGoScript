package main

import (
	"html/template"
	"net/http"
)

// Planet represents information about a planet in the solar system
type Planet struct {
	Name      string
	Diameter  string
	Distance  string
	Mass      string
	OrbitTime string
	ImageURL  string
}

// planetData contains information about the planets of the solar system
var planetData = map[string]Planet{
	"mercury": {
		Name:      "Mercury",
		Diameter:  "4,880 km",
		Distance:  "57.9 million km",
		Mass:      "3.30 x 10^23 kg",
		OrbitTime: "88 days",
		ImageURL:  "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Mercury_in_true_color.jpg/220px-Mercury_in_true_color.jpg",
	},
	"venus": {
		Name:      "Venus",
		Diameter:  "12,104 km",
		Distance:  "108.2 million km",
		Mass:      "4.87 x 10^24 kg",
		OrbitTime: "225 days",
		ImageURL:  "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b2/Venus_2_Approach_Image.jpg/220px-Venus_2_Approach_Image.jpg",
	},
	// Add more planets here...
}

func planetHandler(w http.ResponseWriter, r *http.Request) {
	planetName := r.URL.Path[len("/"):]
	planet, found := planetData[planetName]
	if !found {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.New("planet").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>{{ .Name }}</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					margin: 20px;
					padding: 20px;
				}
				h1 {
					margin-bottom: 10px;
				}
				.planet-info {
					display: flex;
				}
				.planet-img {
					flex: 1;
					padding-right: 20px;
				}
				.planet-data {
					flex: 2;
				}
				.planet-data p {
					margin: 5px 0;
				}
			</style>
		</head>
		<body>
			<h1>{{ .Name }}</h1>
			<div class="planet-info">
				<div class="planet-img">
					<img src="{{ .ImageURL }}" alt="{{ .Name }}" width="300">
				</div>
				<div class="planet-data">
					<p><strong>Diameter:</strong> {{ .Diameter }}</p>
					<p><strong>Distance from Sun:</strong> {{ .Distance }}</p>
					<p><strong>Mass:</strong> {{ .Mass }}</p>
					<p><strong>Orbit Time:</strong> {{ .OrbitTime }}</p>
				</div>
			</div>
		</body>
		</html>
	`))

	tmpl.Execute(w, planet)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/mercury", http.StatusFound)
	})

	http.HandleFunc("/mercury", planetHandler)
	http.HandleFunc("/venus", planetHandler)
	// Add more planet routes here...

	port := "8080" // Change this to your desired port number
	http.ListenAndServe(":"+port, nil)
}
