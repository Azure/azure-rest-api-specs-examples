import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.fluent.models.DatabaseInner;

/** Samples for Databases Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabasesUpdate.json
     */
    /**
     * Sample code: KustoDatabasesUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .update("kustorptest", "kustoclusterrptest4", "KustoDatabase8", new DatabaseInner(), Context.NONE);
    }
}
