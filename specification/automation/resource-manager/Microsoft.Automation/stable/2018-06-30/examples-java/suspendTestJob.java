import com.azure.core.util.Context;

/** Samples for TestJob Suspend. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/suspendTestJob.json
     */
    /**
     * Sample code: Suspend test job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void suspendTestJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .testJobs()
            .suspendWithResponse("mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
