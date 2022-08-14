import com.azure.core.util.Context;

/** Samples for Watcher Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getWatcher.json
     */
    /**
     * Sample code: Get watcher.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getWatcher(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.watchers().getWithResponse("rg", "MyTestAutomationAccount", "MyTestWatcher", Context.NONE);
    }
}
