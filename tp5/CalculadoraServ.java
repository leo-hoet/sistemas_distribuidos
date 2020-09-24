import java.rmi.Naming;
import java.rmi.Remote;
import java.rmi.registry.LocateRegistry;

public class CalculadoraServ {
	
	public CalculadoraServ ()
	{
		try {
		
			//Instancio la clase en memoria, lo devuelvo en una var 
			Calculadora calucladora = new Calculadora();
			
			//Registro calculadora con la URL. Nombrado. Permite invocarlo remotamente
			Naming.rebind("rmi://localhost:1099/Calculadora", (Remote) calucladora);
			System.out.println("Calculadora activa");
			
		} catch (Exception e) {
			System.out.println("Problema: " + e);
		}
	}
	
	public static void main(String[] args) {
		
		CalculadoraServ calcula= new CalculadoraServ();
	
	}

}
