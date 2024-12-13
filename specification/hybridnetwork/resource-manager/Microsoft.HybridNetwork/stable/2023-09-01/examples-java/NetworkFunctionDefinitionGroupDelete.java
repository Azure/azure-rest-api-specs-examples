
/**
 * Samples for NetworkFunctionDefinitionGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionDefinitionGroupDelete.json
     */
    /**
     * Sample code: Delete a network function group resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        deleteANetworkFunctionGroupResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionGroups().delete("rg", "TestPublisher",
            "TestNetworkFunctionDefinitionGroupName", com.azure.core.util.Context.NONE);
    }
}
