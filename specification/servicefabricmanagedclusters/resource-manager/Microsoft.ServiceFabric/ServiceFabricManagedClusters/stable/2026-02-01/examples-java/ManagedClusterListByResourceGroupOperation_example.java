
/**
 * Samples for ManagedClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedClusterListByResourceGroupOperation_example.json
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
