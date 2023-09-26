/** Samples for SourceControlConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/ListSourceControlConfiguration.json
     */
    /**
     * Sample code: List Source Control Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void listSourceControlConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .sourceControlConfigurations()
            .list("rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", com.azure.core.util.Context.NONE);
    }
}
