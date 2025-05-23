
import com.azure.resourcemanager.automation.models.TestJobCreateParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for TestJob Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createTestJob.json
     */
    /**
     * Sample code: Create test job.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void createTestJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.testJobs().createWithResponse("mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial",
            new TestJobCreateParameters()
                .withParameters(mapOf("key01", "fakeTokenPlaceholder", "key02", "fakeTokenPlaceholder")).withRunOn(""),
            com.azure.core.util.Context.NONE);
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
