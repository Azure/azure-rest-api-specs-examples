/** Samples for FluxConfigOperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/GetFluxConfigurationAsyncOperationStatus.json
     */
    /**
     * Sample code: FluxConfigurationAsyncOperationStatus Get.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void fluxConfigurationAsyncOperationStatusGet(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .fluxConfigOperationStatus()
            .getWithResponse(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "srs-fluxconfig",
                "99999999-9999-9999-9999-999999999999",
                com.azure.core.util.Context.NONE);
    }
}
