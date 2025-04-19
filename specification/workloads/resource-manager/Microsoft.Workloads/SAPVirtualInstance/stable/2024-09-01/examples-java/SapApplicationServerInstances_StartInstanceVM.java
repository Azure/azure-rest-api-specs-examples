
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StartRequest;

/**
 * Samples for SapApplicationServerInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_StartInstanceVM.json
     */
    /**
     * Sample code: Start Virtual Machine and the SAP Application Server Instance on it.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startVirtualMachineAndTheSAPApplicationServerInstanceOnIt(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapApplicationServerInstances().start("test-rg", "X00", "app01", new StartRequest().withStartVm(true),
            com.azure.core.util.Context.NONE);
    }
}
