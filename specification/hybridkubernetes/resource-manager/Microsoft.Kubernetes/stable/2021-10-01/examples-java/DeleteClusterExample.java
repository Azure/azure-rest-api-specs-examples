import com.azure.core.util.Context;

/** Samples for ConnectedCluster Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/DeleteClusterExample.json
     */
    /**
     * Sample code: DeleteClusterExample.
     *
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void deleteClusterExample(
        com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().delete("k8sc-rg", "testCluster", Context.NONE);
    }
}
