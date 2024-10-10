
/**
 * Samples for Job Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/suspendJob.json
     */
    /**
     * Sample code: Suspend job.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void suspendJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobs().suspendWithResponse("mygroup", "ContoseAutomationAccount", "foo", null,
            com.azure.core.util.Context.NONE);
    }
}
