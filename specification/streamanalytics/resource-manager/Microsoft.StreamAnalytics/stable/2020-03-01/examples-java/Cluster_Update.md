```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.Cluster;
import com.azure.resourcemanager.streamanalytics.models.ClusterSku;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Cluster_Update.json
     */
    /**
     * Sample code: Update a cluster.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateACluster(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Cluster resource =
            manager.clusters().getByResourceGroupWithResponse("sjrg", "testcluster", Context.NONE).getValue();
        resource.update().withSku(new ClusterSku().withCapacity(96)).apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
