
/**
 * Samples for MongoClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-07-01/MongoClusters_Delete.json
     */
    /**
     * Sample code: Deletes a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        deletesAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().delete("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE);
    }
}
