import com.azure.core.util.Context;

/** Samples for Operations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/ListConnectedClusterOperationsExample.json
     */
    /**
     * Sample code: ListConnectedClusterOperationsExample.
     *
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void listConnectedClusterOperationsExample(
        com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.operations().get(Context.NONE);
    }
}
