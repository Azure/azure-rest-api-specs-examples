
/**
 * Samples for NetworkFunctionDefinitionVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionDefinitionVersionGet.json
     */
    /**
     * Sample code: Get a network function definition version resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getANetworkFunctionDefinitionVersionResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionVersions().getWithResponse("rg", "TestPublisher",
            "TestNetworkFunctionDefinitionGroupName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
