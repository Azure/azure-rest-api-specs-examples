
/**
 * Samples for HorizonDbClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_ListByResourceGroup.json
     */
    /**
     * Sample code: List HorizonDB clusters in a resource group.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listHorizonDBClustersInAResourceGroup(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbClusters().listByResourceGroup("exampleresourcegroup", com.azure.core.util.Context.NONE);
    }
}
