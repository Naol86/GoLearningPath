for _ in range(int(input())):
  a, b = map(int, input().split())
  c, d = map(int, input().split())
  lis = [max(a,c), min(b,d)]
  if lis[0] > lis[1]:
    print(1)
  else:
    ans = lis[1] - lis[0]
    if ans+ 1  <= 0:
      print(1)
    else:
      # ans -= 1
      ans += (a != c)
      ans += (b != d)
      print(ans)