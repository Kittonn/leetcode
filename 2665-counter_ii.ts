interface Counter {
  increment: () => number;
  decrement: () => number;
  reset: () => void;
}

function createCounter(init: number): Counter {
  let count = init;

  const increment = () => ++count;
  const decrement = () => --count;
  const reset = () => (count = init);

  return { increment, decrement, reset };
}
