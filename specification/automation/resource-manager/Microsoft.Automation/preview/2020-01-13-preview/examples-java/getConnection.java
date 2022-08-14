import com.azure.core.util.Context;

/** Samples for Connection Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getConnection.json
     */
    /**
     * Sample code: Get a connection.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getAConnection(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.connections().getWithResponse("rg", "myAutomationAccount28", "myConnection", Context.NONE);
    }
}
