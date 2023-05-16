/** Samples for Jobs StopExecution. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/Job_Stop_Execution.json
     */
    /**
     * Sample code: Terminate a Container Apps Job.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void terminateAContainerAppsJob(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().stopExecution("rg", "testcontainerAppsJob0", "jobExecution1", com.azure.core.util.Context.NONE);
    }
}
