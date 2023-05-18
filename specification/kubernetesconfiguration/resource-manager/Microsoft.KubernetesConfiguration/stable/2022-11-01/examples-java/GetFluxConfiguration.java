/** Samples for FluxConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/GetFluxConfiguration.json
     */
    /**
     * Sample code: Get Flux Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void getFluxConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .fluxConfigurations()
            .getWithResponse(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "srs-fluxconfig",
                com.azure.core.util.Context.NONE);
    }
}
