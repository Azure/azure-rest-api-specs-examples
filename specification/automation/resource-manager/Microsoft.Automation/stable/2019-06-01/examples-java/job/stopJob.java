import com.azure.core.util.Context;

/** Samples for Job Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/stopJob.json
     */
    /**
     * Sample code: Stop job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void stopJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobs().stopWithResponse("mygroup", "ContoseAutomationAccount", "foo", null, Context.NONE);
    }
}
