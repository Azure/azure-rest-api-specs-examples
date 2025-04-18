
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapApplicationServerInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_StopInstanceSoft.json
     */
    /**
     * Sample code: Soft Stop the SAP Application Server Instance.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void softStopTheSAPApplicationServerInstance(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapApplicationServerInstances().stop("test-rg", "X00", "app01",
            new StopRequest().withSoftStopTimeoutSeconds(300L), com.azure.core.util.Context.NONE);
    }
}
