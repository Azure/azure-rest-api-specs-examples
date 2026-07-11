
import com.azure.resourcemanager.azurestackhci.models.CheckUpdatesRequest;

/**
 * Samples for UpdateSummariesOperationGroup CheckUpdates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/UpdateSummaries_CheckUpdates.json
     */
    /**
     * Sample code: Check for updates.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void checkForUpdates(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateSummariesOperationGroups().checkUpdates("testrg", "testcluster", new CheckUpdatesRequest(),
            com.azure.core.util.Context.NONE);
    }
}
