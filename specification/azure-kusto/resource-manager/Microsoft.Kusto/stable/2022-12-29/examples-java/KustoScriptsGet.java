/** Samples for Scripts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoScriptsGet.json
     */
    /**
     * Sample code: KustoScriptsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .scripts()
            .getWithResponse(
                "kustorptest", "kustoCluster", "Kustodatabase8", "kustoScript", com.azure.core.util.Context.NONE);
    }
}
