/**
 * @param {string} val
 * @return {Object}
 */
var expect = (exceptValue) => {
  return {
    toBe: (value) => {
      if (exceptValue === value) {
        return true;
      }
      throw new Error('Not Equal')
    },
    notToBe: (value) => {
      if (exceptValue !== value) {
        return true;
      }
      throw new Error('Equal')
    },
  };
};

console.log(expect(6).toBe(val));
