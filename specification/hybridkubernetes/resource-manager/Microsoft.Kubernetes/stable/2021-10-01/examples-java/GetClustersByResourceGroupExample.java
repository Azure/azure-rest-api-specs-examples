/** Samples for ConnectedCluster ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/GetClustersByResourceGroupExample.json
     */
    /**
     * Sample code: GetClustersExample.
     *
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void getClustersExample(com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().listByResourceGroup("k8sc-rg", com.azure.core.util.Context.NONE);
    }
}
