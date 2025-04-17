
/**
 * Samples for Jobs ListDetectors.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_ListDetectors.json
     */
    /**
     * Sample code: Get the list of available diagnostic data for a Container App Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getTheListOfAvailableDiagnosticDataForAContainerAppJob(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().listDetectors("mikono-workerapp-test-rg", "mikonojob1", com.azure.core.util.Context.NONE);
    }
}
