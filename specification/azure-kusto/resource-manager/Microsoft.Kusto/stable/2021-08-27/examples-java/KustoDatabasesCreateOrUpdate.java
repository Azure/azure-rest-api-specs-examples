import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.fluent.models.DatabaseInner;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabasesCreateOrUpdate.json
     */
    /**
     * Sample code: Kusto ReadWrite database create or update.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoReadWriteDatabaseCreateOrUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .createOrUpdate(
                "kustorptest",
                "kustoclusterrptest4",
                "KustoDatabase8",
                new DatabaseInner().withLocation("westus"),
                Context.NONE);
    }
}
