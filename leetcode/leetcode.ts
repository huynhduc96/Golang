class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;

  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

function createBinaryTreeFromArray(arr: (number | null)[]): TreeNode | null {
  if (arr.length === 0 || arr[0] === null) return null;

  const root = new TreeNode(arr[0]!);
  const queue: TreeNode[] = [root];
  let i = 1;

  while (i < arr.length) {
    const current = queue.shift()!;

    // Assign left child
    if (i < arr.length && arr[i] !== null) {
      current.left = new TreeNode(arr[i]!);
      queue.push(current.left);
    }
    i++;

    // Assign right child
    if (i < arr.length && arr[i] !== null) {
      current.right = new TreeNode(arr[i]!);
      queue.push(current.right);
    }
    i++;
  }

  return root;
}

/**
Given a root of an N-ary tree, you need to compute the length of the diameter of the tree.

The diameter of an N-ary tree is the length of the longest path between any two nodes in the tree. This path may or may not pass through the root.

(Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value.)

Input: root = [1,null,3,2,4,null,5,6]
Output: 3
Explanation: Diameter is shown in red color.

 */

function diameterOfBinaryTree(root: TreeNode | null): number {
  let diameter = 0;
  function helper(root: TreeNode | null): number {
    if (!root) {
      return 0;
    }

    let left = helper(root.left);
    let right = helper(root.right);

    diameter = Math.max(diameter, left + right);

    return Math.max(left, right) + 1;
  }
  helper(root);

  return diameter;
}

function main() {
  const input = [1, 2, 3, 4, 5];
  const root = createBinaryTreeFromArray(input);
  // console.log(JSON.stringify(root, null, 2));
  const result = diameterOfBinaryTree(root);
  console.log("---> result: ", result);
}
main();
