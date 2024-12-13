
/**
 * Samples for Publishers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/PublisherGet.json
     */
    /**
     * Sample code: Get a publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getAPublisherResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.publishers().getByResourceGroupWithResponse("rg", "TestPublisher", com.azure.core.util.Context.NONE);
    }
}
