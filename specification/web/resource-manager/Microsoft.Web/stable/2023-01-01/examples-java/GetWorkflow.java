
/**
 * Samples for WebApps GetInstanceWorkflowSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/GetWorkflow.json
     */
    /**
     * Sample code: GET a workflow Slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETAWorkflowSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getInstanceWorkflowSlotWithResponse("testrg123",
            "testsite2", "staging", "stateful1", com.azure.core.util.Context.NONE);
    }
}
