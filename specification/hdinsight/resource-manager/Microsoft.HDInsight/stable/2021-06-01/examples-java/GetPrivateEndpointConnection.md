Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Get specific private endpoint connection for a specific HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getSpecificPrivateEndpointConnectionForASpecificHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("rg1", "cluster1", "testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2", Context.NONE);
    }
}
```
