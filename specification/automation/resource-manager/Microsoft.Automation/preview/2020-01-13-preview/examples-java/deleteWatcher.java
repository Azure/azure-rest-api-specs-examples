import com.azure.core.util.Context;

/** Samples for Watcher Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteWatcher.json
     */
    /**
     * Sample code: Delete watcher.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteWatcher(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.watchers().deleteWithResponse("rg", "MyTestAutomationAccount", "MyTestWatcher", Context.NONE);
    }
}
