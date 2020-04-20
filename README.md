# jprimego

## Installation

```
go get -u github.com/jefflee1990710/jprimego
```

## Usage

```
p := FastGeneratePrime(1024)
isPrime := p.ProbablyPrime(40)
if !isPrime {
    t.Error("FastGeneratePrime generate non prime number")
}
```

Where p is the generated prime number with 1024 bit length