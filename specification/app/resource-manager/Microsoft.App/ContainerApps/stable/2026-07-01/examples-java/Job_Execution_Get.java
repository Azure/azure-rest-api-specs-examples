
/**
 * Samples for ResourceProvider JobExecution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job_Execution_Get.json
     */
    /**
     * Sample code: Get a single Job Execution.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getASingleJobExecution(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.resourceProviders().jobExecutionWithResponse("rg", "testcontainerAppsJob0", "jobExecution1",
            com.azure.core.util.Context.NONE);
    }
}
