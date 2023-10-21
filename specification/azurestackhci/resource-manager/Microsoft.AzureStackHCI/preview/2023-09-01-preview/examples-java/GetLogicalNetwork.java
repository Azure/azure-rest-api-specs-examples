/** Samples for LogicalNetworksOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/GetLogicalNetwork.json
     */
    /**
     * Sample code: GetLogicalNetwork.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getLogicalNetwork(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .logicalNetworksOperations()
            .getByResourceGroupWithResponse("test-rg", "test-lnet", com.azure.core.util.Context.NONE);
    }
}
