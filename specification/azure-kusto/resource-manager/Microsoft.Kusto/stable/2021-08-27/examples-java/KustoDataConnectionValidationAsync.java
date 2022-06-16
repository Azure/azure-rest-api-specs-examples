import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.fluent.models.DataConnectionValidationInner;
import com.azure.resourcemanager.kusto.models.EventHubDataConnection;

/** Samples for DataConnections DataConnectionValidation. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDataConnectionValidationAsync.json
     */
    /**
     * Sample code: KustoDataConnectionValidation.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionValidation(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .dataConnectionValidation(
                "kustorptest",
                "kustoclusterrptest4",
                "KustoDatabase8",
                new DataConnectionValidationInner()
                    .withDataConnectionName("DataConnections8")
                    .withProperties(new EventHubDataConnection()),
                Context.NONE);
    }
}
