
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StartRequest;

/**
 * Samples for SapCentralInstances StartInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapcentralinstances/SAPCentralInstances_StartInstance_WithInfraOperations.json
     */
    /**
     * Sample code: Start the virtual machine(s) and the SAP central services instance on it.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startTheVirtualMachineSAndTheSAPCentralServicesInstanceOnIt(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralInstances().startInstance("test-rg", "X00", "centralServer",
            new StartRequest().withStartVm(true), com.azure.core.util.Context.NONE);
    }
}
