
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapCentralServerInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralInstances_StopInstanceVM.json
     */
    /**
     * Sample code: Stop the SAP Central Services Instance and its underlying Virtual Machine(s).
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void stopTheSAPCentralServicesInstanceAndItsUnderlyingVirtualMachineS(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralServerInstances().stop("test-rg", "X00", "centralServer",
            new StopRequest().withDeallocateVm(true), com.azure.core.util.Context.NONE);
    }
}
