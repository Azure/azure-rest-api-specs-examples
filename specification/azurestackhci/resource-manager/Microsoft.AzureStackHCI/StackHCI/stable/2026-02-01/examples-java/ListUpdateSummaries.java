
/**
 * Samples for UpdateSummariesOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ListUpdateSummaries.json
     */
    /**
     * Sample code: Get Update summaries under cluster resource.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        getUpdateSummariesUnderClusterResource(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateSummariesOperations().list("testrg", "testcluster", com.azure.core.util.Context.NONE);
    }
}
