const createHelloWorld = (): (() => string) => {
  const helloWorld = (): string => "Hello World";
  return helloWorld;
};
