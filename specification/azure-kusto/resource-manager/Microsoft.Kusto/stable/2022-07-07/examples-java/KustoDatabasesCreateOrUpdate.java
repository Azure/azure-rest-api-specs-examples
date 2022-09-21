import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.CallerRole;
import com.azure.resourcemanager.kusto.models.ReadWriteDatabase;
import java.time.Duration;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoDatabasesCreateOrUpdate.json
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
                "kustoCluster",
                "KustoDatabase8",
                new ReadWriteDatabase().withLocation("westus").withSoftDeletePeriod(Duration.parse("P1D")),
                CallerRole.ADMIN,
                Context.NONE);
    }
}
