#include <cstdio>
#include <iostream>
#include <cstring>
#include <string>
using namespace std;

int T;
string s;
int vis[15];

int main() {
	scanf("%d",&T);
	while(T--) {
		memset(vis,0,sizeof(vis));
		cin >> s;
		int len=s.length();
		for (int i=0;i<len;i++) {
			if (s[i]=='A') vis[1]++;
			else if (s[i]=='0') vis[10]++;
			else if (s[i]=='J') vis[11]++;
			else if (s[i]=='Q') vis[12]++;
			else if (s[i]=='K') vis[13]++;
			else vis[s[i]-'0']++;
		}
		int s1=0,s2=0,s3=0,s4=0,s5=0,s6=0;
		for (int i=1;i<=13;i++) {
			if (vis[i]==4) s4++,s3++,s2++,s1++;
			else if (vis[i]==3) s3++,s2++,s1++;
			else if (vis[i]==2) s2++,s1++;
			else if (vis[i]==1) s1++;
		}
		if (s2) s5=(s2-1)*s3;
		vis[14]=vis[1];
		for (int i=3;i<=10;i++) {
			if (vis[i]&&vis[i+1]&&vis[i+2]&&vis[i+3]&&vis[i+4])
				s6++;
		}
		//printf("%d %d %d %d %d %d\n",s1,s2,s3,s4,s5,s6);
		int ans=s1+s2+s3+s4+s5+s6;
		cout << ans << endl;
	}
	return 0;
}