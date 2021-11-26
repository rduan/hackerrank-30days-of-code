def computeDifference(self):
    s = sorted(self.__elements)
    n = len(self.__elements)
    self.maximumDifference = s[n-1] - s[0]
