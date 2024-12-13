
/**
 * Samples for NetworkFunctionDefinitionVersions ListByNetworkFunctionDefinitionGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionDefinitionVersionListByNetworkFunctionDefinitionGroup.json
     */
    /**
     * Sample code: Get Publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getPublisherResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionVersions().listByNetworkFunctionDefinitionGroup("rg", "TestPublisher",
            "TestNetworkFunctionDefinitionGroupNameName", com.azure.core.util.Context.NONE);
    }
}
