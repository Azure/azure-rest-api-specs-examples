
/**
 * Samples for WebApps ListWorkflowsConnectionsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWorkflowsConfigurationConnections_Slot.json
     */
    /**
     * Sample code: List the Instance Workflows Configuration Connections Slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listTheInstanceWorkflowsConfigurationConnectionsSlot(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listWorkflowsConnectionsSlotWithResponse("testrg123", "testsite2",
            "staging", com.azure.core.util.Context.NONE);
    }
}
