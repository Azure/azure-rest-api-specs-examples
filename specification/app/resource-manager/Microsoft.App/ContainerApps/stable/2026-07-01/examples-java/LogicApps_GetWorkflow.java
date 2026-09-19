
/**
 * Samples for LogicApps GetWorkflow.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/LogicApps_GetWorkflow.json
     */
    /**
     * Sample code: GET a workflow.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void gETAWorkflow(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.logicApps().getWorkflowWithResponse("examplerg", "testcontainerApp0", "testcontainerApp0", "stateful1",
            com.azure.core.util.Context.NONE);
    }
}
