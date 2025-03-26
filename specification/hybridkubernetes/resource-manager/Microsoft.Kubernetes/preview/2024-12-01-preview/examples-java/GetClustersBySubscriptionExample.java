
/**
 * Samples for ConnectedCluster List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/
     * GetClustersBySubscriptionExample.json
     */
    /**
     * Sample code: GetClustersBySubscriptionExample.
     * 
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void
        getClustersBySubscriptionExample(com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().list(com.azure.core.util.Context.NONE);
    }
}
