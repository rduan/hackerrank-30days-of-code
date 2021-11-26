class Node:
	def __init__(self, data):
		self.data = data
		self.left = None
		self.right = None

class Solution:
	def insert(self, root, data):
		if root == None:
			return Node(data)
		else:
			if data <= root.data:
				cur = self.insert(root.left, data)
				root.left = cur
			else:
				cur = self.insert(root.right, data)
				root.right = cur
		return root

	def getHeight(self, root):
		if root == None:
			return -1
		else:
			return 1 + max(self.getHeight(root.left), self.getHeight(root.right))

myTree = Solution()

root=myTree.insert(None, 3)

print(myTree.getHeight(root))
