import com.azure.core.util.Context;

/** Samples for Module ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listModulesByAutomationAccount.json
     */
    /**
     * Sample code: List modules by automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listModulesByAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.modules().listByAutomationAccount("rg", "myAutomationAccount33", Context.NONE);
    }
}
