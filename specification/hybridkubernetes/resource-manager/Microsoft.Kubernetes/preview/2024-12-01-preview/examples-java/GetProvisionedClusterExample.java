
/**
 * Samples for ConnectedCluster GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/
     * GetProvisionedClusterExample.json
     */
    /**
     * Sample code: GetProvisionedClusterExample.
     * 
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void
        getProvisionedClusterExample(com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().getByResourceGroupWithResponse("k8sc-rg", "testCluster",
            com.azure.core.util.Context.NONE);
    }
}
