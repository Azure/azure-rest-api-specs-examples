import com.azure.core.util.Context;

/** Samples for Job Resume. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/resumeJob.json
     */
    /**
     * Sample code: Resume job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void resumeJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobs().resumeWithResponse("mygroup", "ContoseAutomationAccount", "foo", null, Context.NONE);
    }
}
