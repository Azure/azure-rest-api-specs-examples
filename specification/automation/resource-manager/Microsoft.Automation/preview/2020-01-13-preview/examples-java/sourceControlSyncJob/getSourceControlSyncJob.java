
import java.util.UUID;

/**
 * Samples for SourceControlSyncJob Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/
     * sourceControlSyncJob/getSourceControlSyncJob.json
     */
    /**
     * Sample code: Get a source control sync job by job id.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void getASourceControlSyncJobByJobId(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.sourceControlSyncJobs().getWithResponse("rg", "myAutomationAccount33", "MySourceControl",
            UUID.fromString("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a9a"), com.azure.core.util.Context.NONE);
    }
}
