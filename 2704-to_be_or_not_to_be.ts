type ToBeOrNotToBe = {
  toBe: (val: any) => boolean;
  notToBe: (val: any) => boolean;
};

function expect(exceptValue: any): ToBeOrNotToBe {
  const notEqualError = new Error('Not Equal');
  const equalError = new Error('Equal');

  return {
    toBe: (val: any) => {
      if (exceptValue === val) {
        return true;
      }
      throw notEqualError;
    },
    notToBe: (val: any) => {
      if (exceptValue !== val) {
        return true;
      }
      throw equalError;
    },
  };
}
