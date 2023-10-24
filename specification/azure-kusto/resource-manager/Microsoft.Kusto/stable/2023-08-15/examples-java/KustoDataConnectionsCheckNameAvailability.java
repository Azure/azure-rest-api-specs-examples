import com.azure.resourcemanager.kusto.models.DataConnectionCheckNameRequest;

/** Samples for DataConnections CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoDataConnectionsCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                new DataConnectionCheckNameRequest().withName("DataConnections8"),
                com.azure.core.util.Context.NONE);
    }
}
