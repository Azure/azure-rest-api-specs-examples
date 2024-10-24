
/**
 * Samples for Jobs StopMultipleExecutions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/Job_Stop_Multiple.json
     */
    /**
     * Sample code: Terminate Multiple Container Apps Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        terminateMultipleContainerAppsJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().stopMultipleExecutions("rg", "testcontainerAppsJob0", com.azure.core.util.Context.NONE);
    }
}
