
/**
 * Samples for MongoClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mongocluster/DocumentDB.MongoCluster.Management/examples/2024-03-01-preview/MongoClusters_Get.json
     */
    /**
     * Sample code: Gets a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void getsAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster",
            com.azure.core.util.Context.NONE);
    }
}
