import com.azure.resourcemanager.automation.models.RunbookAssociationProperty;
import com.azure.resourcemanager.automation.models.ScheduleAssociationProperty;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/** Samples for JobSchedule Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createJobSchedule.json
     */
    /**
     * Sample code: Create a job schedule.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createAJobSchedule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .jobSchedules()
            .define(UUID.fromString("0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc"))
            .withExistingAutomationAccount("rg", "ContoseAutomationAccount")
            .withSchedule(
                new ScheduleAssociationProperty().withName("ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2"))
            .withRunbook(new RunbookAssociationProperty().withName("TestRunbook"))
            .withParameters(mapOf("jobscheduletag01", "jobschedulevalue01", "jobscheduletag02", "jobschedulevalue02"))
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
