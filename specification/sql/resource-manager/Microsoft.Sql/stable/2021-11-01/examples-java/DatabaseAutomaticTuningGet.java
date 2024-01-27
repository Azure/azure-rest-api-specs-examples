
import com.azure.core.util.Context;

/** Samples for DatabaseAutomaticTuning Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseAutomaticTuningGet.json
     */
    /**
     * Sample code: Get a database's automatic tuning settings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADatabaseSAutomaticTuningSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseAutomaticTunings().getWithResponse("default-sql-onebox",
            "testsvr11", "db1", Context.NONE);
    }
}
