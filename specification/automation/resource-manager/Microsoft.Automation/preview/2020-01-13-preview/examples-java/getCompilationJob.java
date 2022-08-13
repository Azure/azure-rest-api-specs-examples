import com.azure.core.util.Context;

/** Samples for DscCompilationJob Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getCompilationJob.json
     */
    /**
     * Sample code: Get a DSC Compilation job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getADSCCompilationJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.dscCompilationJobs().getWithResponse("rg", "myAutomationAccount33", "TestCompilationJob", Context.NONE);
    }
}
