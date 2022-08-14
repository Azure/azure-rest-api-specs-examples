import com.azure.core.util.Context;

/** Samples for Schedule Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteSchedule.json
     */
    /**
     * Sample code: Delete schedule.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteSchedule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.schedules().deleteWithResponse("rg", "myAutomationAccount33", "mySchedule", Context.NONE);
    }
}
