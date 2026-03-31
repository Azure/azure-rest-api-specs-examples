
/**
 * Samples for WebApps ListWorkflows.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWorkflows.json
     */
    /**
     * Sample code: List the workflows.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listTheWorkflows(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listWorkflows("testrg123", "testsite2", com.azure.core.util.Context.NONE);
    }
}
