import com.azure.core.util.Context;

/** Samples for ConnectionType Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getConnectionType.json
     */
    /**
     * Sample code: Get connection type.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getConnectionType(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.connectionTypes().getWithResponse("rg", "myAutomationAccount22", "myCT", Context.NONE);
    }
}
