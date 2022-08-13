import com.azure.core.util.Context;

/** Samples for Watcher ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listWatchersByAutomationAccount.json
     */
    /**
     * Sample code: List watchers by Automation Account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listWatchersByAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.watchers().listByAutomationAccount("rg", "MyTestAutomationAccount", null, Context.NONE);
    }
}
