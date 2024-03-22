
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StartRequest;

/**
 * Samples for SapApplicationServerInstances StartInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapapplicationinstances/SAPApplicationServerInstances_StartInstance_WithInfraOperations.json
     */
    /**
     * Sample code: Start Virtual Machine and the SAP Application Server Instance on it.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startVirtualMachineAndTheSAPApplicationServerInstanceOnIt(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapApplicationServerInstances().startInstance("test-rg", "X00", "app01",
            new StartRequest().withStartVm(true), com.azure.core.util.Context.NONE);
    }
}
