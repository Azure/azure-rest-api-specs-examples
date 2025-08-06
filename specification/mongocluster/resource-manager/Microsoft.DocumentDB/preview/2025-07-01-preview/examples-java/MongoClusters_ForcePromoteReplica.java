
import com.azure.resourcemanager.mongocluster.models.PromoteMode;
import com.azure.resourcemanager.mongocluster.models.PromoteOption;
import com.azure.resourcemanager.mongocluster.models.PromoteReplicaRequest;

/**
 * Samples for MongoClusters Promote.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_ForcePromoteReplica.json
     */
    /**
     * Sample code: Promotes a replica Mongo Cluster resource to a primary role.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void promotesAReplicaMongoClusterResourceToAPrimaryRole(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().promote("TestGroup", "myMongoCluster",
            new PromoteReplicaRequest().withPromoteOption(PromoteOption.FORCED).withMode(PromoteMode.SWITCHOVER),
            com.azure.core.util.Context.NONE);
    }
}
