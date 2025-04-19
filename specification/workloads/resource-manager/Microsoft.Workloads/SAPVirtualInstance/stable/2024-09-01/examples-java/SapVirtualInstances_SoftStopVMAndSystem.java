
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapVirtualInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_SoftStopVMAndSystem.json
     */
    /**
     * Sample code: Soft Stop the virtual machine(s) and the SAP system on it.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void softStopTheVirtualMachineSAndTheSAPSystemOnIt(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().stop("test-rg", "X00",
            new StopRequest().withSoftStopTimeoutSeconds(300L).withDeallocateVm(true),
            com.azure.core.util.Context.NONE);
    }
}
