const createCounter = (initialValue: number): (() => number) => {
  let counter = initialValue;

  const incrementCounter = (): number => {
    return counter++;
  };

  return incrementCounter;
};
