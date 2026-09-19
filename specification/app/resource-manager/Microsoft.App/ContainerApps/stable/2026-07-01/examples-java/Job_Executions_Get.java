
/**
 * Samples for JobsExecutions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job_Executions_Get.json
     */
    /**
     * Sample code: Get a Container Apps Job Executions.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getAContainerAppsJobExecutions(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobsExecutions().list("rg", "testcontainerAppsJob0", null, com.azure.core.util.Context.NONE);
    }
}
