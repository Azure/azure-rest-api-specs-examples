```java
import com.azure.core.util.Context;

/** Samples for ConnectedCluster GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/GetClusterExample.json
     */
    /**
     * Sample code: GetClusterExample.
     *
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void getClusterExample(com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().getByResourceGroupWithResponse("k8sc-rg", "testCluster", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hybridkubernetes_1.0.0-beta.2/sdk/hybridkubernetes/azure-resourcemanager-hybridkubernetes/README.md) on how to add the SDK to your project and authenticate.
