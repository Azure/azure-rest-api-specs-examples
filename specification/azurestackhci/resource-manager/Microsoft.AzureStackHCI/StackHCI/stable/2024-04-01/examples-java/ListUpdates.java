
/**
 * Samples for Updates List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListUpdates.json
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
