import com.azure.resourcemanager.kusto.models.Script;

/** Samples for Scripts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoScriptsUpdate.json
     */
    /**
     * Sample code: KustoScriptsUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        Script resource =
            manager
                .scripts()
                .getWithResponse(
                    "kustorptest", "kustoCluster", "KustoDatabase8", "kustoScript", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withScriptUrl("https://mysa.blob.core.windows.net/container/script.txt")
            .withForceUpdateTag("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe")
            .withContinueOnErrors(true)
            .apply();
    }
}
