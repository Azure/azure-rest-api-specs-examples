
/**
 * Samples for Clusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ClusterListByResourceGroupOperation_example.json
     */
    /**
     * Sample code: List cluster by resource group.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void
        listClusterByResourceGroup(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusters().listByResourceGroup("resRg", com.azure.core.util.Context.NONE);
    }
}
