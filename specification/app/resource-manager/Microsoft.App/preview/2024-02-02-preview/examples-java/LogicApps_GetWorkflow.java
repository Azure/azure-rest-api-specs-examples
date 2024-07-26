
/**
 * Samples for LogicApps GetWorkflow.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/LogicApps_GetWorkflow.json
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
