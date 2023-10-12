import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateMode;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/ServerCreateGeoRestore.json
     */
    /**
     * Sample code: Create a database as a geo-restore in geo-paired location.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createADatabaseAsAGeoRestoreInGeoPairedLocation(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .servers()
            .define("pgtestsvc5geo")
            .withRegion("eastus")
            .withExistingResourceGroup("testrg")
            .withSourceServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/sourcepgservername")
            .withPointInTimeUtc(OffsetDateTime.parse("2021-06-27T00:04:59.4078005+00:00"))
            .withCreateMode(CreateMode.GEO_RESTORE)
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
