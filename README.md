<h6 align="center"> <img src="https://github.com/npsolver/Mongolang/blob/master/public/logo.png" width=400 height=400 /></h6>
A tool to generate Golang code from MongoDB queries. 

## Motivation behind the Project
While working as a Software Engineering intern, I had to write Golang code to query a MongoDB database. I quickly found the process tedious, as it required constantly referring to the MongoDB Go Driver documentation just to construct queries.

I searched online for a tool that could help automate this, but couldn’t find anything that did what I needed. That led me to wonder: why not build one myself?

At its core, the idea was simple. The input would be a MongoDB query—which already resembles a programming language—and the output would be equivalent Golang code. This immediately reminded me of concepts I learned in CS241: Foundations of Sequential Programs at the University of Waterloo, where we built a compiler that translated a subset of C++ into Assembly code.

Seeing the similarity, I decided to approach this project in the same way: by treating MongoDB queries as a source language and generating Golang code as the target. The process I followed is described below:

## Process
The overall compiling process involves 4 steps:

### Step 1: Scanning
First and foremost, we are to scan the MongoDB query and get a sequence of tokens representing the query. This is implemented using a Deterministic Finite Automata (DFA). More info [here](https://github.com/npsolver/Mongolang/tree/master/core/scanner). 

### Step 2: Parsing
Next up, using the parsed tokens, we use an extended form of DFA to generate a parsetree. This is done to make sure the query follows a specific context-free grammar. More info [here](https://github.com/npsolver/Mongolang/tree/master/core/parser). 

### Step 3: Context-Sensitive Analysis
In this stage, we check all of the conditions which scanning or parsing couldn't handle. An example might be to check whether the field mentioned in the query is actually in the database being called. As of now, since the project is still in its primary stages, we do not need to implement anything for this step. The tool just converts any given MongoDB query into Golang code.

### Step 4: Code Generation and Optimization
We can in a pretty straightforward fashion generate the Golang code using the parsetree. For now, we are only generating the filter to be passed in a code that uses MongoDB's Golang driver. In future, we hope to implement full code generation so that users can just define the queries and directly work with the generated code.

## Example
For an input MongoDB query,
```
db.users.find({
    $or: [
        { age: 
            { $lt: 18 } 
        }, 
        { age: 
            { $gt: 65 } }
    ]
})
```
we get the Golang code,
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

## Future Development
We aim to provide fully working Golang code that one can use out of the box, instead of just providing the filter above.
