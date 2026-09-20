
/**
 * Samples for LogicApps ListWorkflows.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/LogicApps_ListWorkflows.json
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
