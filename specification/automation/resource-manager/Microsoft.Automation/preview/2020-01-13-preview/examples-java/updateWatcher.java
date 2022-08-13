import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.Watcher;

/** Samples for Watcher Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateWatcher.json
     */
    /**
     * Sample code: Update watcher.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateWatcher(com.azure.resourcemanager.automation.AutomationManager manager) {
        Watcher resource =
            manager
                .watchers()
                .getWithResponse("rg", "MyTestAutomationAccount", "MyTestWatcher", Context.NONE)
                .getValue();
        resource.update().withName("MyTestWatcher").withExecutionFrequencyInSeconds(600L).apply();
    }
}
