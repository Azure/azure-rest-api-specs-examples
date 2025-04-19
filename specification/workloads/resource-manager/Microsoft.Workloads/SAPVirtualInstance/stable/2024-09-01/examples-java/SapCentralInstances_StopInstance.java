
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapCentralServerInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralInstances_StopInstance.json
     */
    /**
     * Sample code: Stop the SAP Central Services Instance.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void stopTheSAPCentralServicesInstance(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralServerInstances().stop("test-rg", "X00", "centralServer",
            new StopRequest().withSoftStopTimeoutSeconds(1200L), com.azure.core.util.Context.NONE);
    }
}
