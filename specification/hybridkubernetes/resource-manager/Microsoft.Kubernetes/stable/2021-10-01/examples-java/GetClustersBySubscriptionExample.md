Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hybridkubernetes_1.0.0-beta.2/sdk/hybridkubernetes/azure-resourcemanager-hybridkubernetes/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ConnectedCluster List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/GetClustersBySubscriptionExample.json
     */
    /**
     * Sample code: GetClustersExample.
     *
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void getClustersExample(com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().list(Context.NONE);
    }
}
```
