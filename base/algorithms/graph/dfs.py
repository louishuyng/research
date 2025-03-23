class Solution:
    # This is DFS when we need to use stack to implement it.
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        n = len(rooms)
        if n <= 1:
            return True
        stack = [rooms[0]]
        visited = {0}

        while stack:
            room = stack.pop()
            for key in room:
                if key in visited:
                    continue
                stack.append(rooms[key])
                visited.add(key)
        if len(visited) == n:
            return True
        else:
            return False

    # This is DFS when we need to use recursion to implement it.
    def findCircleNum(self, isConnected: List[List[int]]) -> int:
        def dfs(node, d,v):
            v[node] = True
            for i in d[node]:
                if not v[i]:
                    dfs(i, d, v)
        
        n = len(isConnected)

        d = {i: [] for i in range(n)}

        for i in range(n):
            for j in range(n):
                if isConnected[i][j]==1 and i!=j:               
                    d[i].append(j)
                    d[j].append(i)
        v=[False]*n
        ans = 0
        for i in range(n):
            if not v[i]:
                ans += 1
                dfs(i, d,v)
        return ans
        
if __name__ == '__main__':
    rooms = [[1],[2],[3],[]]
    isConnected = [[1,1,0],[1,1,0],[0,0,1]]
    print(Solution().canVisitAllRooms(rooms))
    print(Solution().findCircleNum(isConnected))
