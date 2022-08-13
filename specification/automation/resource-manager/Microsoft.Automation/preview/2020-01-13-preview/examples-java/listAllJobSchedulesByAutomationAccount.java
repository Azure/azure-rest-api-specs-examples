import com.azure.core.util.Context;

/** Samples for JobSchedule ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listAllJobSchedulesByAutomationAccount.json
     */
    /**
     * Sample code: List all job schedules by automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listAllJobSchedulesByAutomationAccount(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobSchedules().listByAutomationAccount("rg", "ContoseAutomationAccount", null, Context.NONE);
    }
}
