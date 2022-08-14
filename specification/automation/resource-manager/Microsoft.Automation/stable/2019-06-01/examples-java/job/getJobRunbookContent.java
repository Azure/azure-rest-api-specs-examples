import com.azure.core.util.Context;

/** Samples for Job GetRunbookContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/getJobRunbookContent.json
     */
    /**
     * Sample code: Get Job Runbook Content.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getJobRunbookContent(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobs().getRunbookContentWithResponse("mygroup", "ContoseAutomationAccount", "foo", null, Context.NONE);
    }
}
