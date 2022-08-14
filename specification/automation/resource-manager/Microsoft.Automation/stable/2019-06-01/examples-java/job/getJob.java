import com.azure.core.util.Context;

/** Samples for Job Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/getJob.json
     */
    /**
     * Sample code: Get job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobs().getWithResponse("mygroup", "ContoseAutomationAccount", "foo", null, Context.NONE);
    }
}
