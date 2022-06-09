```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetAllPrivateEndpointConnectionsInCluster.json
     */
    /**
     * Sample code: Get all private endpoint connections for a specific HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllPrivateEndpointConnectionsForASpecificHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.privateEndpointConnections().listByCluster("rg1", "cluster1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.
