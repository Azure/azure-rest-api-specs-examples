import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.ReadOnlyFollowingDatabase;
import java.time.Duration;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDatabaseReadonlyUpdate.json
     */
    /**
     * Sample code: Kusto ReadOnly database update.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoReadOnlyDatabaseUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .createOrUpdate(
                "kustorptest",
                "kustoCluster",
                "kustoReadOnlyDatabase",
                new ReadOnlyFollowingDatabase().withLocation("westus").withHotCachePeriod(Duration.parse("P1D")),
                Context.NONE);
    }
}
