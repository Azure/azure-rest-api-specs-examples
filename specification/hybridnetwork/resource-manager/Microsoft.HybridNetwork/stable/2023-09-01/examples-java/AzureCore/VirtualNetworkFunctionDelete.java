
/**
 * Samples for NetworkFunctions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureCore/
     * VirtualNetworkFunctionDelete.json
     */
    /**
     * Sample code: Delete virtual network function resource on AzureCore.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteVirtualNetworkFunctionResourceOnAzureCore(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().delete("rg", "testNf", com.azure.core.util.Context.NONE);
    }
}
