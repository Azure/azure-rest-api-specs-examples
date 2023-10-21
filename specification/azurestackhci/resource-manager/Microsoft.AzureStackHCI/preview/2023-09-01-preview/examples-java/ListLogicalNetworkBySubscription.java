/** Samples for LogicalNetworksOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListLogicalNetworkBySubscription.json
     */
    /**
     * Sample code: ListLogicalNetworkBySubscription.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listLogicalNetworkBySubscription(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.logicalNetworksOperations().list(com.azure.core.util.Context.NONE);
    }
}
