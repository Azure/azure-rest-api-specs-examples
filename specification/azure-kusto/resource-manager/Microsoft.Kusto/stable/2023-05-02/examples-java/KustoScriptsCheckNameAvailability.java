import com.azure.resourcemanager.kusto.models.ScriptCheckNameRequest;

/** Samples for Scripts CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoScriptsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoScriptsCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .scripts()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoCluster",
                "db",
                new ScriptCheckNameRequest().withName("kustoScriptName1"),
                com.azure.core.util.Context.NONE);
    }
}
