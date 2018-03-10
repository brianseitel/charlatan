# Charlatan

Charlatan is a fake data generator for Go that leans on struct tags to help build dummy data.

## Getting Started

Getting started is simple:

```
go get github.com/brianseitel/charlatan
```

To integrate into your app, define a struct and use the appropriate struct tags. Here's quick example:

```golang
type User struct {
    ID int `charlatan:"id"`
    Name string `charlatan:"name"`
    Email string `charlatan:"email"`
}

func main() {
	u := charlatan.New().Generate(&User{})
	fmt.Printf("%# v\n", u)
}
```

## Fake Data Types

### Address
| Type | Tag | Example | Example |
|---|---|---|---|
| address | `address` | `123 Main St, Anywhere, MA 00250` | `string` |
| city | `city` | `New York` | `string` |
| latitude | `latitude` | `72.5315` | `float64` |
| longitude | `longitude` | `-102.555` | `float64` |
| streetName | `streetName` | `Main St` | `string` |
| zip | `zip` | `94210` | `string` |
| state | `state` | `California` | `string` |
| stateAbbr | `stateAbbr` | `CA` | `string` |

### Base
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| digit | `digit` | `1` | `int64` |
| number | `number` | `130` | `int64` |
| boolean | `boolean` | `true` | `bool` |
| float | `float` | `35.316262` | `float64` |
| price | `price` | `2.95` | `float64` |
| age | `age` | `17` | `int64` |
| id | `id` | `8814021` | `int64` |
| letter | `letter` | `c` | `string` |
| url | `url` | `https://www.example.com/food` | `string` |

### Crypto
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| sha1 | `sha1` | ~` | `string` |
| sha256 | `sha256` | ~ | `string` |
| md5 | `md5` | ~ | `string` |

### Date/Time
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| unixTimestamp | `unixTimestamp` | ~` | `int64` |
| datetime | `datetime` | `2018-03-04T15:01:02Z` | `string` |
| date | `date` | `2018-03-04` | `string` |
| time | `time` | `12:19:32` | `string` |
| dayOfWeek | `dayOfWeek` | `Monday` | `string` |
| dayOfMonth | `dayOfMonth` | `28` | `int64` |
| month | `month` | `1` | `int64` |
| monthName | `monthName` | `January` | `string` |

### Email
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| email | `email` | `cheese@burgler.com` | `string` |


### Internet
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| ipv4 | `ipv4` | `205.161.119.32` | `string` |
| ipv6 | `ipv6` | `ff7f:de06:dbea:a7ff:77ef:eae0"` | `string` |


### Person
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| firstName | `firstName` | `Jane` | `string` |
| lastName | `lastName` | `Doe` | `string` |
| fullName | `fullName` | `Jane Doe` | `string` |


### Phone
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| phone | `phone` | `555-876-5309` | `string` |


### Barcodes
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| ean | `ean` | `436967481997487` | `string` |


### UUIDs
| Type | Tag | Example | Expected Type |
|---|---|---|---|
| uuid | `uuid` | `bfab5831-4a48-468a-b0ca-358e2378abbd` | `string` |


### Words
| Type | Tag | Description | Expected Type |
|---|---|---|---|
| word | `word` | `cheese` | `string` |
| words | `words` | `[ice, cream, sandwich]` | `[]string` |

## Custom Tags

If the built-in rules above aren't good enough for you, you can add your own! It's as easy as 1, 2, 3:

### One: Define a Function

The signature for the function is `func() interface{}`:

```golang
func generateColor() interface{} {
	charset := "1234567890ABCDEF"

	var letters []byte
	for i := 0; i < 6; i++ {
		idx := rand.Intn(len(charset))
		letters = append(letters, charset[idx])
	}

	return fmt.Sprintf("#%s", string(letters))
}
```

### Two: Add it to your Charlatan instance

The signature for adding a custom tag is `AddCustomTag(name string, callback func() interface{})`

```golang
c := charlatan.New()
c.AddCustomTag("hexcolor", generateColor)
```

### Three: Use it in your struct!

Now you can add it as a struct tag in your struct:

```golang
type App struct {
	BackgroundColor string `charlatan:"hexcolor"`
	ForegroundColor string `charlatan:"hexcolor"`
}
```

## Examples

You can run any of the examples by running the following:

```
go run --tags=examples /examples/{example}/main.go
```

## Contributing

Fork it and submit a PR!

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/brianseitel/charlatan/tags). 

## Authors

* **Brian Seitel** - *Initial work* - [Brian Seitel](https://github.com/PurpleBooth)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
