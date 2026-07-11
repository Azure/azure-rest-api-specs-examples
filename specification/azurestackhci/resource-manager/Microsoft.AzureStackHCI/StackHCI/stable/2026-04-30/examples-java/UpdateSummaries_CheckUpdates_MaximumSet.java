
import com.azure.resourcemanager.azurestackhci.models.CheckUpdatesRequest;

/**
 * Samples for UpdateSummariesOperationGroup CheckUpdates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/UpdateSummaries_CheckUpdates_MaximumSet.json
     */
    /**
     * Sample code: Check for specific update by name.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        checkForSpecificUpdateByName(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateSummariesOperationGroups().checkUpdates("testrg", "testcluster",
            new CheckUpdatesRequest().withUpdateName("Microsoft4.2203.2.32"), com.azure.core.util.Context.NONE);
    }
}
