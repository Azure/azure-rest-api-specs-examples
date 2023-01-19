import com.azure.resourcemanager.mariadb.models.ServerPropertiesForReplica;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreateReplicaMode.json
     */
    /**
     * Sample code: Create a replica server.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void createAReplicaServer(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .servers()
            .define("targetserver")
            .withRegion("westus")
            .withExistingResourceGroup("TargetResourceGroup")
            .withProperties(
                new ServerPropertiesForReplica()
                    .withSourceServerId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/MasterResourceGroup/providers/Microsoft.DBforMariaDB/servers/masterserver"))
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
