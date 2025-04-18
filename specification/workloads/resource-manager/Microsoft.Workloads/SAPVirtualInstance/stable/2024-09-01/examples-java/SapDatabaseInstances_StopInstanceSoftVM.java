
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapDatabaseInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapDatabaseInstances_StopInstanceSoftVM.json
     */
    /**
     * Sample code: Soft Stop the database instance of the SAP system and the underlying Virtual Machine(s).
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void softStopTheDatabaseInstanceOfTheSAPSystemAndTheUnderlyingVirtualMachineS(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapDatabaseInstances().stop("test-rg", "X00", "db0",
            new StopRequest().withSoftStopTimeoutSeconds(300L).withDeallocateVm(true),
            com.azure.core.util.Context.NONE);
    }
}
