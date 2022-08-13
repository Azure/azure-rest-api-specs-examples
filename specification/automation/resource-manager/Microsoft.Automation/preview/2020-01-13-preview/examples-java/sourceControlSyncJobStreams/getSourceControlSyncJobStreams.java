import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for SourceControlSyncJobStreams ListBySyncJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControlSyncJobStreams/getSourceControlSyncJobStreams.json
     */
    /**
     * Sample code: Get a list of sync job streams identified by sync job id.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getAListOfSyncJobStreamsIdentifiedBySyncJobId(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .sourceControlSyncJobStreams()
            .listBySyncJob(
                "rg",
                "myAutomationAccount33",
                "MySourceControl",
                UUID.fromString("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a2b"),
                null,
                Context.NONE);
    }
}
