
/**
 * Samples for UpdateSummariesOperationGroup CheckHealth.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/UpdateSummaries_CheckHealth.json
     */
    /**
     * Sample code: Check health of UpdateSummaries.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        checkHealthOfUpdateSummaries(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateSummariesOperationGroups().checkHealth("testrg", "testcluster", com.azure.core.util.Context.NONE);
    }
}
