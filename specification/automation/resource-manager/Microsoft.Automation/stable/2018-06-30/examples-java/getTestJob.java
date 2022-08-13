import com.azure.core.util.Context;

/** Samples for TestJob Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getTestJob.json
     */
    /**
     * Sample code: Get test job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getTestJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.testJobs().getWithResponse("mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
