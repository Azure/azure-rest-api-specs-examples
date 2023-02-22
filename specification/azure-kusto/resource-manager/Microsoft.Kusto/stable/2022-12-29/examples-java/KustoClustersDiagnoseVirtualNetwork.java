/** Samples for Clusters DiagnoseVirtualNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoClustersDiagnoseVirtualNetwork.json
     */
    /**
     * Sample code: KustoClusterDiagnoseVirtualNetwork.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterDiagnoseVirtualNetwork(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().diagnoseVirtualNetwork("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
