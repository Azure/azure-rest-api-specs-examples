
/**
 * Samples for Publishers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * PublisherListByResourceGroup.json
     */
    /**
     * Sample code: List all publisher resources in a resource group.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllPublisherResourcesInAResourceGroup(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.publishers().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
