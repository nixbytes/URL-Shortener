# URL-Shortener

exercise status: released
Exercise details

The goal of this exercise is to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs we would look for any incoming web requests with the path /dogs and redirect them.

To complete this exercises you will need to implement the stubbed out methods in handler.go. There are a good bit of comments explaining what each method should do, and there is also a main/main.go source file that uses the package to help you test your code and get an idea of what your program should be doing.

I suggest first commenting out all of the code in main.go related to the YAMLHandler function and focusing on implementing the MapHandler function first.

Once you have that working, focus on parsing the YAML using the gopkg.in/yaml.v2 package. Note: You will need to go get this package if you don’t have it already.

After you get the YAML parsing down, try to convert the data into a map and then use the MapHandler to finish the YAMLHandler implementation. Eg you might end up with some code like this:

```go
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
  parsedYaml, err := parseYAML(yaml)
  if err != nil {
    return nil, err
  }
  pathMap := buildMap(parsedYaml)
  return MapHandler(pathMap, fallback), nil
}
```
But in order for this to work you will need to create functions like parseYAML and buildMap on your own. This should give you ample experience working with YAML data.
