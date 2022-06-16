import com.azure.core.util.Context;

/** Samples for Scripts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoScriptsDelete.json
     */
    /**
     * Sample code: KustoScriptsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.scripts().delete("kustorptest", "kustoCluster", "KustoDatabase8", "kustoScript", Context.NONE);
    }
}
