for _ in range(int(input())):
  n, k = map(int, input().split())
  nums = [int(i) for i in input().split()]
  nums.sort(reverse=True)
  alice = 0
  bob = 0
  pre_alice = 0
  for i in range(n):
    if i % 2 == 0:
      alice += nums[i]
      pre_alice = nums[i]
    else:
      diff = pre_alice - nums[i]
      if diff == 0:
        bob += nums[i]
      elif k > diff:
        bob += pre_alice
        k -= diff
      elif k > 0:
        bob += nums[i] + k
        k = 0
      else:
        bob += nums[i]
  print(alice - bob)