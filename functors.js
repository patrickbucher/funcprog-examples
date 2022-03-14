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
