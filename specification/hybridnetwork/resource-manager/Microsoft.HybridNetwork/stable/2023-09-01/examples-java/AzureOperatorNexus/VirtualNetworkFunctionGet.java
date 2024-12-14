
/**
 * Samples for NetworkFunctions GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * AzureOperatorNexus/VirtualNetworkFunctionGet.json
     */
    /**
     * Sample code: Get virtual network function resource on AzureOperatorNexus.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getVirtualNetworkFunctionResourceOnAzureOperatorNexus(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().getByResourceGroupWithResponse("rg", "testNf", com.azure.core.util.Context.NONE);
    }
}
