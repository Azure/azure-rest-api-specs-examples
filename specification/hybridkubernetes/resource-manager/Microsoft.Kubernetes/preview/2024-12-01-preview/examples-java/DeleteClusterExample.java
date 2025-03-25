
/**
 * Samples for ConnectedCluster Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/
     * DeleteClusterExample.json
     */
    /**
     * Sample code: DeleteClusterExample.
     * 
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void
        deleteClusterExample(com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().delete("k8sc-rg", "testCluster", com.azure.core.util.Context.NONE);
    }
}
