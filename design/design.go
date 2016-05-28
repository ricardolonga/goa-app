package design                                     // The convention consists of naming the design
                                                   // package "design"
import (
        . "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("clients", func() {                     // API defines the microservice endpoint and
        Title("The virtual clients catalog")        // other global properties. There should be one
        Description("A simple goa service")        // and exactly one API definition appearing in
        Scheme("http")                             // the design.
        Host("localhost:8080")
})

var _ = Resource("client", func() {                // Resources group related API endpoints
        BasePath("/clients")                       // together. They map to REST resources for REST
        DefaultMedia(ClientMedia)                  // services.

        Action("show", func() {                    // Actions define a single API endpoint together
                Description("Get client by id")    // with its path, parameters (both path
                Routing(GET("/:clientId"))         // parameters and querystring values) and payload
                Params(func() {                    // (shape of the request body).
                        Param("clientId", Integer, "Client Id")
                })
                Response(OK)                       // Responses define the shape and status code
                Response(NotFound)                 // of HTTP responses.
        })
})

// BottleMedia defines the media type used to render bottles.
var ClientMedia = MediaType("application/vnd.goa.example.client+json", func() {
        Description("A client")
        Attributes(func() {                         // Attributes define the media type shape.
                Attribute("id", Integer, "Unique client Id")
                Attribute("href", String, "API href for making requests on the client")
                Attribute("name", String, "Name of client")
                Required("id", "href", "name")
        })
        View("default", func() {                    // View defines a rendering of the media type.
                Attribute("id")                     // Media types may have multiple views and must
                Attribute("href")                   // have a "default" view.
                Attribute("name")
        })
})
