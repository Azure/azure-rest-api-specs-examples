
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StartRequest;

/**
 * Samples for SapDatabaseInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapDatabaseInstances_StartInstanceVM.json
     */
    /**
     * Sample code: Start Virtual Machine and the database instance of the SAP system on it.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startVirtualMachineAndTheDatabaseInstanceOfTheSAPSystemOnIt(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapDatabaseInstances().start("test-rg", "X00", "db0", new StartRequest().withStartVm(true),
            com.azure.core.util.Context.NONE);
    }
}
