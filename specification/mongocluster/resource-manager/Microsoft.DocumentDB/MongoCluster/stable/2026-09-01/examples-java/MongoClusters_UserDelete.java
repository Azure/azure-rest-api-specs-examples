
/**
 * Samples for Users Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_UserDelete.json
     */
    /**
     * Sample code: Deletes a user on a DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        deletesAUserOnADocumentDBClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.users().delete("TestGroup", "myMongoCluster", "uuuuuuuu-uuuu-uuuu-uuuu-uuuuuuuuuuuu",
            com.azure.core.util.Context.NONE);
    }
}
