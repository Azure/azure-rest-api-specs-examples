import com.azure.core.util.Context;

/** Samples for TestJobStreams ListByTestJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/listTestJobStreamsByJob.json
     */
    /**
     * Sample code: List job streams by job name.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listJobStreamsByJobName(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .testJobStreams()
            .listByTestJob("mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", null, Context.NONE);
    }
}
