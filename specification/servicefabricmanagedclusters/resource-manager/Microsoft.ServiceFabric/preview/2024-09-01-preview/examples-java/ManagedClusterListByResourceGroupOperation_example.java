
/**
 * Samples for ManagedClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/ManagedClusterListByResourceGroupOperation_example.json
     */
    /**
     * Sample code: List cluster by resource group.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listClusterByResourceGroup(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusters().listByResourceGroup("resRg", com.azure.core.util.Context.NONE);
    }
}
