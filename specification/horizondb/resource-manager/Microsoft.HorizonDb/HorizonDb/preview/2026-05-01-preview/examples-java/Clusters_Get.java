
/**
 * Samples for HorizonDbClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_Get.json
     */
    /**
     * Sample code: Get a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBCluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbClusters().getByResourceGroupWithResponse("exampleresourcegroup", "examplecluster",
            com.azure.core.util.Context.NONE);
    }
}
