#include "calcular.h"
int main (int argc, char **argv){
    char *host;
    CLIENT *sv;
    int res,sum;
    float mlt;
    host = argv[1];
    sv = clnt_create(host, calcular, UNO, "tcp");
    if (sv == NULL) {
        clnt_pcreateerror(host);
        exit(0);
    }      
    sum = *(sumar_1(5, 2, sv));
    res = *(restar_1(5,5,sv));
    mlt = *(multiplicar_1(10,20,sv));
    if ((&sum == NULL)|| (&res == NULL) || (&mlt == NULL)) {
        clnt_perror(sv, "call failed");
        exit(0);
    }
    printf("El resultado de la suma es : %d\n",sum);
    printf("El resultado de la resta es: %d\n", res);
    printf("El resultado de la mlt es: %f.3\n",mlt);
    clnt_destroy(sv);
    exit(0);
}
