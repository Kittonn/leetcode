function lengthOfLastWord(s: string): number {
  const word = s.trim();

  const lastSpaceIndex = word.lastIndexOf(" ");

  return word.length - lastSpaceIndex - 1;
}
