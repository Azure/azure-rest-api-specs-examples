
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StartRequest;

/**
 * Samples for SapCentralServerInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralInstances_StartInstanceVM.json
     */
    /**
     * Sample code: Start the virtual machine(s) and the SAP central services instance on it.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startTheVirtualMachineSAndTheSAPCentralServicesInstanceOnIt(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralServerInstances().start("test-rg", "X00", "centralServer",
            new StartRequest().withStartVm(true), com.azure.core.util.Context.NONE);
    }
}
