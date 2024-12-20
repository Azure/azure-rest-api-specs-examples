
import java.util.UUID;

/**
 * Samples for SourceControlSyncJob Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/
     * sourceControlSyncJob/createSourceControlSyncJob.json
     */
    /**
     * Sample code: Create or update a source control sync job.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void
        createOrUpdateASourceControlSyncJob(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.sourceControlSyncJobs().define(UUID.fromString("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a9a"))
            .withExistingSourceControl("rg", "myAutomationAccount33", "MySourceControl")
            .withCommitId("9de0980bfb45026a3d97a1b0522d98a9f604226e").create();
    }
}
