
import com.azure.resourcemanager.hdinsight.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.hdinsight.models.PrivateLinkServiceConnectionStatus;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * ApprovePrivateEndpointConnection.json
     */
    /**
     * Sample code: Approve a private endpoint connection manually.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void
        approveAPrivateEndpointConnectionManually(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.privateEndpointConnections().define("testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2")
            .withExistingCluster("rg1", "cluster1")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                    .withDescription("update it from pending to approved.").withActionsRequired("None"))
            .create();
    }
}
