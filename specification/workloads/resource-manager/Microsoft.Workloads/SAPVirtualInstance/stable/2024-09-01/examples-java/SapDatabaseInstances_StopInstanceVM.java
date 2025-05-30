
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapDatabaseInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapDatabaseInstances_StopInstanceVM.json
     */
    /**
     * Sample code: Stop the database instance of the SAP system and the underlying Virtual Machine(s).
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void stopTheDatabaseInstanceOfTheSAPSystemAndTheUnderlyingVirtualMachineS(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapDatabaseInstances().stop("test-rg", "X00", "db0",
            new StopRequest().withSoftStopTimeoutSeconds(0L).withDeallocateVm(true), com.azure.core.util.Context.NONE);
    }
}
