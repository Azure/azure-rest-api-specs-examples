
/**
 * Samples for Users ListByMongoCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/MongoClusters_UserList.json
     */
    /**
     * Sample code: List the users on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        listTheUsersOnAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.users().listByMongoCluster("TestGroup", "myMongoCluster", com.azure.core.util.Context.NONE);
    }
}
