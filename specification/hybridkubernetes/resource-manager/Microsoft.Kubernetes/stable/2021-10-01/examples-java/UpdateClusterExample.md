```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.hybridkubernetes.models.ConnectedCluster;
import java.util.HashMap;
import java.util.Map;

/** Samples for ConnectedCluster Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/UpdateClusterExample.json
     */
    /**
     * Sample code: UpdateClusterExample.
     *
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void updateClusterExample(
        com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        ConnectedCluster resource =
            manager
                .connectedClusters()
                .getByResourceGroupWithResponse("k8sc-rg", "testCluster", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hybridkubernetes_1.0.0-beta.2/sdk/hybridkubernetes/azure-resourcemanager-hybridkubernetes/README.md) on how to add the SDK to your project and authenticate.
