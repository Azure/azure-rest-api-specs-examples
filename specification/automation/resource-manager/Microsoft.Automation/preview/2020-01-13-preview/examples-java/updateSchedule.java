import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.Schedule;

/** Samples for Schedule Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateSchedule.json
     */
    /**
     * Sample code: Update a schedule.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateASchedule(com.azure.resourcemanager.automation.AutomationManager manager) {
        Schedule resource =
            manager.schedules().getWithResponse("rg", "myAutomationAccount33", "mySchedule", Context.NONE).getValue();
        resource
            .update()
            .withName("mySchedule")
            .withDescription("my updated description of schedule goes here")
            .withIsEnabled(false)
            .apply();
    }
}
