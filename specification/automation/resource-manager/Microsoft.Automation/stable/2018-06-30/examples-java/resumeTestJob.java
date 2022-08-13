import com.azure.core.util.Context;

/** Samples for TestJob Resume. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/resumeTestJob.json
     */
    /**
     * Sample code: Resume test job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void resumeTestJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .testJobs()
            .resumeWithResponse("mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
