import com.azure.resourcemanager.kusto.models.ReadWriteDatabase;
import java.time.Duration;

/** Samples for Databases Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDatabasesUpdate.json
     */
    /**
     * Sample code: KustoDatabasesUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .update(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                new ReadWriteDatabase().withHotCachePeriod(Duration.parse("P1D")),
                null,
                com.azure.core.util.Context.NONE);
    }
}
