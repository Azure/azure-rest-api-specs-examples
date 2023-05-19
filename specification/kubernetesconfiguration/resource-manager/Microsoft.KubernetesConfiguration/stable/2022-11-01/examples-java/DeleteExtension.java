/** Samples for Extensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/DeleteExtension.json
     */
    /**
     * Sample code: Delete Extension.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void deleteExtension(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .extensions()
            .delete(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "ClusterMonitor",
                null,
                com.azure.core.util.Context.NONE);
    }
}
