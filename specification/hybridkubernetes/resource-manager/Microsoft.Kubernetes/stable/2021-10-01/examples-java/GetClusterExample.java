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
