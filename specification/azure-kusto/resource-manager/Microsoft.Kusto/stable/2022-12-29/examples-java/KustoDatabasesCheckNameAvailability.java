import com.azure.resourcemanager.kusto.models.CheckNameRequest;
import com.azure.resourcemanager.kusto.models.Type;

/** Samples for Databases CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoDatabasesCheckNameAvailability.json
     */
    /**
     * Sample code: KustoDatabasesCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoCluster",
                new CheckNameRequest().withName("database1").withType(Type.MICROSOFT_KUSTO_CLUSTERS_DATABASES),
                com.azure.core.util.Context.NONE);
    }
}
