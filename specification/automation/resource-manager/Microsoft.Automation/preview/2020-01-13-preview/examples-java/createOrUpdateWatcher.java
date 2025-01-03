
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Watcher CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/
     * createOrUpdateWatcher.json
     */
    /**
     * Sample code: Create or update watcher.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateWatcher(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.watchers().define("MyTestWatcher").withExistingAutomationAccount("rg", "MyTestAutomationAccount")
            .withTags(mapOf()).withExecutionFrequencyInSeconds(60L).withScriptName("MyTestWatcherRunbook")
            .withScriptRunOn("MyTestHybridWorkerGroup").withDescription("This is a test watcher.").create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
