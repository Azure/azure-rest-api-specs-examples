/** Samples for Extensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/GetExtension.json
     */
    /**
     * Sample code: Get Extension.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void getExtension(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .extensions()
            .getWithResponse(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "ClusterMonitor",
                com.azure.core.util.Context.NONE);
    }
}
