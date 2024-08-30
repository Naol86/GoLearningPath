for _ in range(int(input())):
  x = int(input())
  nums = [int(i) for i in input().split()]
  if x > 2:
    print("NO")
  else:
    if abs(nums[-1] - nums[0]) > 1:
      print("YES")
    else:
      print("NO")