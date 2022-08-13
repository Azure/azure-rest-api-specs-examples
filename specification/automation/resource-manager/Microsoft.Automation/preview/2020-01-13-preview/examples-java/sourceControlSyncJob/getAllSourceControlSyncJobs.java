import com.azure.core.util.Context;

/** Samples for SourceControlSyncJob ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControlSyncJob/getAllSourceControlSyncJobs.json
     */
    /**
     * Sample code: Get a list of source control sync jobs.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getAListOfSourceControlSyncJobs(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .sourceControlSyncJobs()
            .listByAutomationAccount("rg", "myAutomationAccount33", "MySourceControl", null, Context.NONE);
    }
}
