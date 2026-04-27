
/**
 * Samples for Updates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/DeleteUpdates.json
     */
    /**
     * Sample code: Delete an Update.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteAnUpdate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updates().delete("testrg", "testcluster", "Microsoft4.2203.2.32", com.azure.core.util.Context.NONE);
    }
}
