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
        manager.connectedClusters().list(com.azure.core.util.Context.NONE);
    }
}
