/** Samples for FluxConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/DeleteFluxConfiguration.json
     */
    /**
     * Sample code: Delete Flux Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void deleteFluxConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .fluxConfigurations()
            .delete(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "srs-fluxconfig",
                null,
                com.azure.core.util.Context.NONE);
    }
}
