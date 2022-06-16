import com.azure.core.util.Context;

/** Samples for SourceControlConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/GetSourceControlConfiguration.json
     */
    /**
     * Sample code: Get Source Control Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void getSourceControlConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .sourceControlConfigurations()
            .getWithResponse(
                "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "SRS_GitHubConfig", Context.NONE);
    }
}
