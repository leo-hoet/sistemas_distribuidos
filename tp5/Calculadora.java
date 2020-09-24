import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;

public class Calculadora extends UnicastRemoteObject implements CalculadoraRemota{

	Calculadora() throws RemoteException{
		super();
	}
	
	
	public int sumar(int a, int b) throws RemoteException
	{
		return a+b;
	}
	
	public int restar (int a, int b) throws RemoteException
	{
		
			return a-b;
		
	}
	
	public int multipilicar(int a, int b) throws RemoteException
	{
		return a*b;
		
	}
	
	public int dividir(int a, int b) throws RemoteException
	{
		if(b!=0)
		{
			return a/b;
		}
		else {System.out.println("La división por 0 no existe");
		return 0;
		}
	}
	
}
