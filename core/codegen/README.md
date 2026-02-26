# Code generation

In this step we transform the parsetree we got from the parsing step to actual Golang driver code. We can also generate a template code to work out of the box without the user having to manually set it up. But for now, we are just providing the filter itself corresponding to MongoDB query.

At each level of the parsetree, we have one of set, list, field, operator, etc. It is generally easy to directly convert a given type to the corresponding golang code using the parsetree and recursion. We do not go super into detail here, but it is fairly easy to understand how it might work looking at the example below. We also format the output before returning it.

## Example

For a parsetree
```
START
|---BOF BOF
|---COMMAND
|   |---DBNAME
|   |   |---ID db
|   |---DOT .
|   |---COLLECTIONNAME
|   |   |---ID users
|   |---DOT .
|   |---COLLECTIONCOMMAND
|       |---ID find
|       |---FIRSTBRACKETLEFT (
|       |---OPTIONS
|       |   |---OPTION
|       |       |---SET
|       |           |---SECONDBRACKETLEFT {
|       |           |---FIELDS
|       |           |   |---FIELD
|       |           |       |---FIELDNAME
|       |           |       |   |---ID age
|       |           |       |---COLON :
|       |           |       |---VALUE
|       |           |           |---INT 18
|       |           |---SECONDBRACKETRIGHT }
|       |---FIRSTBRACKETRIGHT )
|---EOF EOF
```
we get the generated code:
```
bson.M{
        "age": 18
}
```
For a different and more complicated parsetree, we will get the following output (not including the partree here since it is a bit long and confusing):
```
bson.M{
	"$or": bson.A{
		bson.M{
			"age": bson.M{
				"$lt": 18
			}
		},
		bson.M{
			"age": bson.M{
				"$gt": 65
			}
		}
	}
}
```
