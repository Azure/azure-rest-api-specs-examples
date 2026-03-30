
/**
 * Samples for WebApps GetInstanceWorkflowSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWorkflow_Slot.json
     */
    /**
     * Sample code: GET a workflow Slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void gETAWorkflowSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getInstanceWorkflowSlotWithResponse("testrg123", "testsite2", "staging",
            "stateful1", com.azure.core.util.Context.NONE);
    }
}
