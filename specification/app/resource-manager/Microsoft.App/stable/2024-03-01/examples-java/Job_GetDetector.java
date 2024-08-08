
/**
 * Samples for Jobs GetDetector.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Job_GetDetector.json
     */
    /**
     * Sample code: Get diagnostic data for a Container App Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getDiagnosticDataForAContainerAppJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().getDetectorWithResponse("mikono-workerapp-test-rg", "mikonojob1", "containerappjobnetworkIO",
            com.azure.core.util.Context.NONE);
    }
}
