Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.Extension;

/** Samples for Extensions Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateExtension.json
     */
    /**
     * Sample code: Create a monitoring extension on Hadoop Linux cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void createAMonitoringExtensionOnHadoopLinuxCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .extensions()
            .create(
                "rg1",
                "cluster1",
                "clustermonitoring",
                new Extension().withWorkspaceId("a2090ead-8c9f-4fba-b70e-533e3e003163").withPrimaryKey("**********"),
                Context.NONE);
    }
}
```
