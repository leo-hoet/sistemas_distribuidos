import java.rmi.Naming;

public class CalculadoraCliente {
	public static void main(String[] args)
	{
		try {
			//Busco la calculadora y llamo a los metodos como si fuera metodo local
			Calculadora calculadora = (Calculadora) Naming.lookup("rmi://localhost:1010/Calculadora");
			
			//Ejecuto todos los metodos de la calc que quiera
			System.out.println("\nBusqueda en Server Remoto Exitosa");
			System.out.println(calculadora.sumar(5, 10));
			
		} catch (Exception e) {
			System.out.println("Excepcion" + e);
		}
	}

}
