import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateMode;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/ServerCreateReviveDropped.json
     */
    /**
     * Sample code: ServerCreateReviveDropped.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverCreateReviveDropped(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .servers()
            .define("pgtestsvc5-rev")
            .withRegion("westus")
            .withExistingResourceGroup("testrg")
            .withSourceServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc5")
            .withPointInTimeUtc(OffsetDateTime.parse("2023-04-27T00:04:59.4078005+00:00"))
            .withCreateMode(CreateMode.REVIVE_DROPPED)
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
