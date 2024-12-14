
/**
 * Samples for NetworkFunctions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * AzureOperatorNexus/VirtualNetworkFunctionDelete.json
     */
    /**
     * Sample code: Delete virtual network function resource on AzureOperatorNexus.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteVirtualNetworkFunctionResourceOnAzureOperatorNexus(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().delete("rg", "testNf", com.azure.core.util.Context.NONE);
    }
}
