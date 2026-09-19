
/**
 * Samples for Jobs StopMultipleExecutions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job_Stop_Multiple.json
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
