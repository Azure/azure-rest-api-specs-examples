import com.azure.resourcemanager.mysqlflexibleserver.models.CreateMode;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2022-09-30-preview/examples/ServerCreateReplica.json
     */
    /**
     * Sample code: Create a replica server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createAReplicaServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .servers()
            .define("replica-server")
            .withRegion("SoutheastAsia")
            .withExistingResourceGroup("testgr")
            .withCreateMode(CreateMode.REPLICA)
            .withSourceServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testgr/providers/Microsoft.DBforMySQL/flexibleServers/source-server")
            .create();
    }

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
