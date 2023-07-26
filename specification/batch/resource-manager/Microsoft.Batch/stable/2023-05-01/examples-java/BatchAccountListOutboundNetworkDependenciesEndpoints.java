/** Samples for BatchAccount ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/BatchAccountListOutboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: ListOutboundNetworkDependencies.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listOutboundNetworkDependencies(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .batchAccounts()
            .listOutboundNetworkDependenciesEndpoints(
                "default-azurebatch-japaneast", "sampleacct", com.azure.core.util.Context.NONE);
    }
}
