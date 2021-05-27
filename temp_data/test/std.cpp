#include <cstdio>
#include <cmath>
using namespace std;

struct Point {
	double x, y;
	Point(double x=0, double y=0):x(x),y(y) {}
};

double Odistance(Point a, Point b) {
	return sqrt((a.x-b.x)*(a.x-b.x)+(a.y-b.y)*(a.y-b.y));
}

int main() {
	int T;
	double a,b,c,d;
	scanf("%d",&T);
	while(T--) {
		scanf("%lf%lf%lf%lf",&a,&b,&c,&d);
		Point A(a,-b),B(c,d);
		printf("%.2lf\n",Odistance(A,B));
	}
	return 0;
}