
/**
 * Samples for WorkflowTriggerHistories List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * WorkflowTriggerHistories_List.json
     */
    /**
     * Sample code: List a workflow trigger history.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAWorkflowTriggerHistory(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowTriggerHistories().list("testResourceGroup", "test-name",
            "testWorkflowName", "testTriggerName", null, null, com.azure.core.util.Context.NONE);
    }
}
