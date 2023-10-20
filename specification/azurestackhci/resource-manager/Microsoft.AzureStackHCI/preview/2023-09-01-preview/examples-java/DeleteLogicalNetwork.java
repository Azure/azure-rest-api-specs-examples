/** Samples for LogicalNetworksOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteLogicalNetwork.json
     */
    /**
     * Sample code: DeleteLogicalNetwork.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteLogicalNetwork(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.logicalNetworksOperations().delete("test-rg", "test-lnet", com.azure.core.util.Context.NONE);
    }
}
