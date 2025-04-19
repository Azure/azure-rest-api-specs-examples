
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapVirtualInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_SoftStop.json
     */
    /**
     * Sample code: Soft Stop of SapVirtualInstances_Stop.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void softStopOfSapVirtualInstancesStop(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().stop("test-rg", "X00", new StopRequest().withSoftStopTimeoutSeconds(300L),
            com.azure.core.util.Context.NONE);
    }
}
