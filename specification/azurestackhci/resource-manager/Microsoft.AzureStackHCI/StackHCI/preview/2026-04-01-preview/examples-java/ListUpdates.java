
/**
 * Samples for Updates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ListUpdates.json
     */
    /**
     * Sample code: List available updates.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listAvailableUpdates(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updates().list("testrg", "testcluster", com.azure.core.util.Context.NONE);
    }
}
