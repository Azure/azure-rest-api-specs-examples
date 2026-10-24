
/**
 * Samples for Clusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_GetAfterPoolCreationOnEmptyCluster.json
     */
    /**
     * Sample code: Get a cluster after pool creation while it had no existing pools.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAClusterAfterPoolCreationWhileItHadNoExistingPools(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().getByResourceGroupWithResponse("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
