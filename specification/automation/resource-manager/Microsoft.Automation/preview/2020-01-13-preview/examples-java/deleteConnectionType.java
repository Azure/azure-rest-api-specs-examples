import com.azure.core.util.Context;

/** Samples for ConnectionType Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteConnectionType.json
     */
    /**
     * Sample code: Delete an existing connection type.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteAnExistingConnectionType(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.connectionTypes().deleteWithResponse("rg", "myAutomationAccount22", "myCT", Context.NONE);
    }
}
