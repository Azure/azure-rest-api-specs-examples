import com.azure.core.util.Context;

/** Samples for Scripts ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoScriptsListByDatabase.json
     */
    /**
     * Sample code: KustoScriptsList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.scripts().listByDatabase("kustorptest", "kustoclusterrptest4", "Kustodatabase8", Context.NONE);
    }
}
