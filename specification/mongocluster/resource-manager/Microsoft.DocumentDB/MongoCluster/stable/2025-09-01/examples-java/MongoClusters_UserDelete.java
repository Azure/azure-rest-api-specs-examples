
/**
 * Samples for Users Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/MongoClusters_UserDelete.json
     */
    /**
     * Sample code: Deletes a user on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        deletesAUserOnAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.users().delete("TestGroup", "myMongoCluster", "uuuuuuuu-uuuu-uuuu-uuuu-uuuuuuuuuuuu",
            com.azure.core.util.Context.NONE);
    }
}
