
/**
 * Samples for WebApps ListInstanceWorkflowsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWorkflows_Slot.json
     */
    /**
     * Sample code: List the workflows Slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listTheWorkflowsSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listInstanceWorkflowsSlot("testrg123", "testsite2", "staging",
            com.azure.core.util.Context.NONE);
    }
}
