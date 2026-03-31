
/**
 * Samples for WebApps GetWorkflow.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWorkflow.json
     */
    /**
     * Sample code: GET a workflow.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void gETAWorkflow(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getWorkflowWithResponse("testrg123", "testsite2", "stateful1",
            com.azure.core.util.Context.NONE);
    }
}
