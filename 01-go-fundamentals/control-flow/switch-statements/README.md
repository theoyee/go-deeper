# Switch Statements

## What I Learned

A `switch` statement compares one value against multiple possible values.

---

## Basic Syntax

```go id="e8n4r2"
switch value {
case option1:
	// code
case option2:
	// code
default:
	// fallback
}
```

---

## Example

```go id="1j7p5m"
day := "Monday"

switch day {
case "Monday":
	fmt.Println("Start of the week")
case "Friday":
	fmt.Println("Almost the weekend")
default:
	fmt.Println("Regular day")
}
```

The program compares `day` against each case.

The matching case executes.

---

## Important Go Behavior

Go does not automatically fall through to the next case.

A `break` is not required after each case.

Only the matching case runs.

---

## `default`

The `default` case runs when none of the other cases match.

It is similar to the `else` branch of an `if` statement.

---

## Switch vs Else If

`switch` is useful when comparing one value against multiple known options.

Instead of:

```go id="6x9qk3"
if day == "Monday" {
} else if day == "Friday" {
}
```

A `switch` can make the logic cleaner and easier to read.

---

## Mental Model

> Use `switch` when asking: "Which option is this value?"

---

## What I Practiced

I created a `switch` statement that handled different commands.

I also used `default` to handle unknown commands.

---

## Questions

- What does a `switch` statement do?
- What is a `case`?
- What does `default` do?
- Does Go require `break` after each case?
- When is `switch` better than multiple `else if` statements?
