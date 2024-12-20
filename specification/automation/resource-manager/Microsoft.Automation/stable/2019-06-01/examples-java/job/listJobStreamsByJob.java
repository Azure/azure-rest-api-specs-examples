
/**
 * Samples for JobStream ListByJob.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/listJobStreamsByJob
     * .json
     */
    /**
     * Sample code: List job streams by job name.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void listJobStreamsByJobName(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobStreams().listByJob("mygroup", "ContoseAutomationAccount", "foo", null, null,
            com.azure.core.util.Context.NONE);
    }
}
