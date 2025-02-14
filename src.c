#include<stdio.h>
#include<math.h>

int main(){
	float a,b,c,num,ans1,ans2;
	scanf("%f %f %f",&a,&b,&c);
	num=b*b-4*a*c;
	if(num<0){
		printf("None");
	}
	else{
		ans1=(-1*b+sqrt(num))/(2*a);
		ans2=(-1*b-sqrt(num))/(2*a);
		if(ans1>ans2){
			printf("%.3f %.3f",ans2,ans1);
		}
		else if(ans2>ans1){
			printf("%.3f %.3f",ans1,ans2);
		}
		else{
			printf("%.3f",ans1);
		}
	}
}