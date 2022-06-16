import com.azure.core.util.Context;

/** Samples for Clusters DiagnoseVirtualNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClustersDiagnoseVirtualNetwork.json
     */
    /**
     * Sample code: KustoClusterDiagnoseVirtualNetwork.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterDiagnoseVirtualNetwork(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().diagnoseVirtualNetwork("kustorptest", "kustoclusterrptest4", Context.NONE);
    }
}
