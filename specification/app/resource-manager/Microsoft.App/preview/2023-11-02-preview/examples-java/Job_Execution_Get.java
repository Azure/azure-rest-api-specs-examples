
/**
 * Samples for ResourceProvider JobExecution.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Job_Execution_Get.json
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
