Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.avs.models.Cluster;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Clusters_Update.json
     */
    /**
     * Sample code: Clusters_Update.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void clustersUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        Cluster resource = manager.clusters().getWithResponse("group1", "cloud1", "cluster1", Context.NONE).getValue();
        resource.update().withClusterSize(4).apply();
    }
}
```
