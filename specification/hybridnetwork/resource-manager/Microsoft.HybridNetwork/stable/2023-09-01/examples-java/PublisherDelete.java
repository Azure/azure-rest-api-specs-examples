
/**
 * Samples for Publishers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/PublisherDelete.
     * json
     */
    /**
     * Sample code: Delete a publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteAPublisherResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.publishers().delete("rg", "TestPublisher", com.azure.core.util.Context.NONE);
    }
}
