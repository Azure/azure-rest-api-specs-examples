import com.azure.core.util.Context;

/** Samples for Scripts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoScriptsDelete.json
     */
    /**
     * Sample code: KustoScriptsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.scripts().delete("kustorptest", "kustoclusterrptest4", "KustoDatabase8", "kustoScript1", Context.NONE);
    }
}
