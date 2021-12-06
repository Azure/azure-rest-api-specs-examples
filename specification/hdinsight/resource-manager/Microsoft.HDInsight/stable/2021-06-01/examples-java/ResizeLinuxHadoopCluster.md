Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.ClusterResizeParameters;
import com.azure.resourcemanager.hdinsight.models.RoleName;

/** Samples for Clusters Resize. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/ResizeLinuxHadoopCluster.json
     */
    /**
     * Sample code: Resize the worker nodes for a Hadoop on Linux cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void resizeTheWorkerNodesForAHadoopOnLinuxCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .clusters()
            .resize(
                "rg1",
                "cluster1",
                RoleName.WORKERNODE,
                new ClusterResizeParameters().withTargetInstanceCount(10),
                Context.NONE);
    }
}
```
