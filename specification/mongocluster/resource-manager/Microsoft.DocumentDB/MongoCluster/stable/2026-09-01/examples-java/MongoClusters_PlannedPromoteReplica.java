
import com.azure.resourcemanager.mongocluster.models.PromoteMode;
import com.azure.resourcemanager.mongocluster.models.PromoteOption;
import com.azure.resourcemanager.mongocluster.models.PromoteReplicaRequest;

/**
 * Samples for MongoClusters Promote.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_PlannedPromoteReplica.json
     */
    /**
     * Sample code: Promotes a replica DocumentDB cluster resource to a primary role, waiting for the replica to catch
     * up before promoting.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        promotesAReplicaDocumentDBClusterResourceToAPrimaryRoleWaitingForTheReplicaToCatchUpBeforePromoting(
            com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().promote("TestGroup", "myMongoCluster",
            new PromoteReplicaRequest().withPromoteOption(PromoteOption.PLANNED).withMode(PromoteMode.SWITCHOVER),
            com.azure.core.util.Context.NONE);
    }
}
