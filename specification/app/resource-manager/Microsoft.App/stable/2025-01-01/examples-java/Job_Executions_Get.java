
/**
 * Samples for JobsExecutions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_Executions_Get.json
     */
    /**
     * Sample code: Get a Container Apps Job Executions.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getAContainerAppsJobExecutions(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobsExecutions().list("rg", "testcontainerappsjob0", null, com.azure.core.util.Context.NONE);
    }
}
