/** Samples for LogicalNetworksOperation ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListLogicalNetworkByResourceGroup.json
     */
    /**
     * Sample code: ListLogicalNetworkByResourceGroup.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listLogicalNetworkByResourceGroup(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.logicalNetworksOperations().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
