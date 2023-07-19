/** Samples for Clusters ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoOutboundNetworkDependenciesList.json
     */
    /**
     * Sample code: Get Kusto cluster outbound network dependencies.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void getKustoClusterOutboundNetworkDependencies(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusters()
            .listOutboundNetworkDependenciesEndpoints("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
