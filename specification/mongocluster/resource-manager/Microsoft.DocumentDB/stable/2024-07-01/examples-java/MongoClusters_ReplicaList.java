
/**
 * Samples for Replicas ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-07-01/MongoClusters_ReplicaList.json
     */
    /**
     * Sample code: List the replicas linked to a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listTheReplicasLinkedToAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.replicas().listByParent("TestGroup", "myMongoCluster", com.azure.core.util.Context.NONE);
    }
}
