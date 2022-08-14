import com.azure.core.util.Context;

/** Samples for TestJob Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/stopTestJob.json
     */
    /**
     * Sample code: Stop test job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void stopTestJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.testJobs().stopWithResponse("mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
