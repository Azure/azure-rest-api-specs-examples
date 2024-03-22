
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapCentralInstances StopInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapcentralinstances/SAPCentralInstances_StopInstance_WithInfraOperations.json
     */
    /**
     * Sample code: Stop the SAP Central Services Instance and its underlying Virtual Machine(s).
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void stopTheSAPCentralServicesInstanceAndItsUnderlyingVirtualMachineS(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralInstances().stopInstance("test-rg", "X00", "centralServer",
            new StopRequest().withDeallocateVm(true), com.azure.core.util.Context.NONE);
    }
}
