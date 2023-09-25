/** Samples for FluxConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/ListFluxConfigurations.json
     */
    /**
     * Sample code: List Flux Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void listFluxConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .fluxConfigurations()
            .list("rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", com.azure.core.util.Context.NONE);
    }
}
