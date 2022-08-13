import com.azure.core.util.Context;

/** Samples for ConnectionType ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getAllConnectionTypes_First100.json
     */
    /**
     * Sample code: Get connection types, first 100.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getConnectionTypesFirst100(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.connectionTypes().listByAutomationAccount("rg", "myAutomationAccount25", Context.NONE);
    }
}
