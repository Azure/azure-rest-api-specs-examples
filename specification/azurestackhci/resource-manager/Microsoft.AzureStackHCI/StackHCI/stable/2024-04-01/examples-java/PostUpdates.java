
/**
 * Samples for Updates Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * PostUpdates.json
     */
    /**
     * Sample code: List available updates.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listAvailableUpdates(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updates().post("testrg", "testcluster", "Microsoft4.2203.2.32", com.azure.core.util.Context.NONE);
    }
}
