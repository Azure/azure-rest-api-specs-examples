
/**
 * Samples for Jobs StopExecution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job_Stop_Execution.json
     */
    /**
     * Sample code: Terminate a Container Apps Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        terminateAContainerAppsJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().stopExecution("rg", "testcontainerAppsJob0", "jobExecution1", com.azure.core.util.Context.NONE);
    }
}
