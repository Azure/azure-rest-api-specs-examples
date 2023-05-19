/** Samples for SourceControlConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/DeleteSourceControlConfiguration.json
     */
    /**
     * Sample code: Delete Source Control Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void deleteSourceControlConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .sourceControlConfigurations()
            .delete(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "SRS_GitHubConfig",
                com.azure.core.util.Context.NONE);
    }
}
