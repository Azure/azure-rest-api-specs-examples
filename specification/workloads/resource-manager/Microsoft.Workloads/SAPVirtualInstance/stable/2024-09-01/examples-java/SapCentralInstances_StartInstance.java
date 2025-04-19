
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StartRequest;

/**
 * Samples for SapCentralServerInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralInstances_StartInstance.json
     */
    /**
     * Sample code: Start the SAP Central Services Instance.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startTheSAPCentralServicesInstance(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralServerInstances().start("test-rg", "X00", "centralServer", new StartRequest(),
            com.azure.core.util.Context.NONE);
    }
}
