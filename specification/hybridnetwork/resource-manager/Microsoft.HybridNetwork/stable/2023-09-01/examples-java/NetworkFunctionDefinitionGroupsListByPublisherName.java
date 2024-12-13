
/**
 * Samples for NetworkFunctionDefinitionGroups ListByPublisher.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionDefinitionGroupsListByPublisherName.json
     */
    /**
     * Sample code: Get networkFunctionDefinition groups under publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getNetworkFunctionDefinitionGroupsUnderPublisherResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionGroups().listByPublisher("rg", "TestPublisher",
            com.azure.core.util.Context.NONE);
    }
}
