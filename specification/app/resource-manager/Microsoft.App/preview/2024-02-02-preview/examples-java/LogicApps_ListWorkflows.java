
/**
 * Samples for LogicApps ListWorkflows.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/LogicApps_ListWorkflows.json
     */
    /**
     * Sample code: List the workflows.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listTheWorkflows(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.logicApps().listWorkflows("examplerg", "testcontainerApp0", "testcontainerApp0",
            com.azure.core.util.Context.NONE);
    }
}
