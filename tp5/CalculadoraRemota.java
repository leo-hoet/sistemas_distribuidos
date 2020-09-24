import java.rmi.Remote;
import java.rmi.RemoteException;

public interface CalculadoraRemota extends java.rmi.Remote{
	
	public int sumar(int a, int b) throws RemoteException;
	
	public int restar (int a, int b) throws RemoteException;
	
	public int multipilicar(int a, int b) throws RemoteException;
	
	public int dividir(int a, int b) throws RemoteException;

}
