import com.azure.core.util.Context;

/** Samples for Watcher Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/stopWatcher.json
     */
    /**
     * Sample code: Start Watcher.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void startWatcher(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.watchers().stopWithResponse("rg", "MyTestAutomationAccount", "MyTestWatcher", Context.NONE);
    }
}
