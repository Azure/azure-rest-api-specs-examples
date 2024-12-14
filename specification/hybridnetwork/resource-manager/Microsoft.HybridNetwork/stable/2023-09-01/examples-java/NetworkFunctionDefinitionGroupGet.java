
/**
 * Samples for NetworkFunctionDefinitionGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionDefinitionGroupGet.json
     */
    /**
     * Sample code: Get a networkFunctionDefinition group resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getANetworkFunctionDefinitionGroupResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionGroups().getWithResponse("rg", "TestPublisher",
            "TestNetworkFunctionDefinitionGroupName", com.azure.core.util.Context.NONE);
    }
}
