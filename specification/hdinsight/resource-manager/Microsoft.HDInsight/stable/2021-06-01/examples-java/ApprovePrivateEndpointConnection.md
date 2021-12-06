Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.hdinsight.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.hdinsight.models.PrivateLinkServiceConnectionStatus;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/ApprovePrivateEndpointConnection.json
     */
    /**
     * Sample code: Approve a private endpoint connection manually.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void approveAPrivateEndpointConnectionManually(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .privateEndpointConnections()
            .define("testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2")
            .withExistingCluster("rg1", "cluster1")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                    .withDescription("update it from pending to approved.")
                    .withActionsRequired("None"))
            .create();
    }
}
```
