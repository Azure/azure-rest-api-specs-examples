
/**
 * Samples for Connection Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/
     * deleteConnection.json
     */
    /**
     * Sample code: Delete an existing connection.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteAnExistingConnection(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.connections().deleteWithResponse("rg", "myAutomationAccount28", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}
