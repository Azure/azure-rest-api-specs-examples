
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateMode;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * ServersCreateReviveDropped.json
     */
    /**
     * Sample code: Create a new server using a backup of a server that was deleted or dropped recently.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createANewServerUsingABackupOfAServerThatWasDeletedOrDroppedRecently(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().define("exampleserver").withRegion("eastus").withExistingResourceGroup("exampleresourcegroup")
            .withSourceServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampledeletedserver")
            .withPointInTimeUtc(OffsetDateTime.parse("2025-06-01T18:30:22.123456Z"))
            .withCreateMode(CreateMode.REVIVE_DROPPED).create();
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
