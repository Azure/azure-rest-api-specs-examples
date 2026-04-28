
/**
 * Samples for UpdateSummariesOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/DeleteUpdateSummaries.json
     */
    /**
     * Sample code: Delete an Update.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteAnUpdate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateSummariesOperations().delete("testrg", "testcluster", com.azure.core.util.Context.NONE);
    }
}
