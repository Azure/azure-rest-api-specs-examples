
/**
 * Samples for Users ListByMongoCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_UserList.json
     */
    /**
     * Sample code: List the users on a DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        listTheUsersOnADocumentDBClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.users().listByMongoCluster("TestGroup", "myMongoCluster", com.azure.core.util.Context.NONE);
    }
}
