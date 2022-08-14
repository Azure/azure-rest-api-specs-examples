import com.azure.core.util.Context;

/** Samples for Schedule Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getSchedule.json
     */
    /**
     * Sample code: Get a schedule.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getASchedule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.schedules().getWithResponse("rg", "myAutomationAccount33", "mySchedule", Context.NONE);
    }
}
