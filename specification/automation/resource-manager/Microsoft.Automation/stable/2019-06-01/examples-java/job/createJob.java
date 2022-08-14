import com.azure.resourcemanager.automation.models.RunbookAssociationProperty;
import java.util.HashMap;
import java.util.Map;

/** Samples for Job Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/createJob.json
     */
    /**
     * Sample code: Create job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .jobs()
            .define("foo")
            .withExistingAutomationAccount("mygroup", "ContoseAutomationAccount")
            .withRunbook(new RunbookAssociationProperty().withName("TestRunbook"))
            .withParameters(mapOf("key01", "value01", "key02", "value02"))
            .withRunOn("")
            .create();
    }

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
