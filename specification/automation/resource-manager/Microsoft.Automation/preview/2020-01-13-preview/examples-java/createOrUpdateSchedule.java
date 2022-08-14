import com.azure.resourcemanager.automation.models.AdvancedSchedule;
import com.azure.resourcemanager.automation.models.ScheduleFrequency;
import java.time.OffsetDateTime;

/** Samples for Schedule CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateSchedule.json
     */
    /**
     * Sample code: Create or update a schedule.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateASchedule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .schedules()
            .define("mySchedule")
            .withExistingAutomationAccount("rg", "myAutomationAccount33")
            .withName("mySchedule")
            .withStartTime(OffsetDateTime.parse("2017-03-27T17:28:57.2494819Z"))
            .withFrequency(ScheduleFrequency.HOUR)
            .withDescription("my description of schedule goes here")
            .withExpiryTime(OffsetDateTime.parse("2017-04-01T17:28:57.2494819Z"))
            .withInterval(1)
            .withAdvancedSchedule(new AdvancedSchedule())
            .create();
    }
}
