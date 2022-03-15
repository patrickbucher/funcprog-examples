function whatAreFunctors() {
    // Functors are things to which map() can be applied.
    const map = (items, f) => {
        const transformed = [];
        for (const item of items) {
            transformed.push(f(item));
        }
        return transformed;
    };

    // An array is a functor.
    const numbersTwice = map([1, 2, 3], x => x * 2);
    console.log(numbersTwice);

    // The corner-case of the empty array is handled without further ado.
    const noNumbersTwice = map([], x => x * 2);
    console.log(noNumbersTwice);

    // The array provides its own map() method:
    const numbersAsString = [1, 2, 3].map(x => x + '');
    console.log(numbersAsString);

    // A promise is a functor, but provides the method then() instead of map().
    const promiseNumber = Promise.resolve(7);
    const promiseTwiceThatNumber = promiseNumber.then(x => x * 2);
    promiseTwiceThatNumber.then(console.log);
}

const numberedLetters = {
    a: 1,
    b: 2,
    c: 3
};

const mapObject = (object, f) => {
    const entries = Object.entries(object);
    const mappedObject = {};
    entries.forEach(([key, value]) => {
        mappedObject[key] = f(value);
    });
    return mappedObject; 
};

const stringNumberedLetters = mapObject(numberedLetters, x => x + '');
console.log(stringNumberedLetters);

class Maybe {
    constructor(value) {
        this.value = value;
    }

    static just(value) {
        if (value === null || value === undefined) {
            throw new Error("Can't construct a value from null/undefined");
        }
        return new Maybe(value);
    }

    static nothing() {
        return new Maybe(null);
    }

    map(f) {
        if (this.value === null) {
            return this;
        }
        return new Maybe(f(this.value));
    }
}

const maybeNumber = Maybe.just(5);
const maybeString = maybeNumber.map(x => x + '');
const maybeAnotherString = Maybe.nothing().map(x => x + '');
