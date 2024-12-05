
import com.azure.resourcemanager.mysqlflexibleserver.models.EnableStatusEnum;
import com.azure.resourcemanager.mysqlflexibleserver.models.Network;
import com.azure.resourcemanager.mysqlflexibleserver.models.Server;
import com.azure.resourcemanager.mysqlflexibleserver.models.Storage;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/ServerUpdate
     * .json
     */
    /**
     * Sample code: Update a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void updateAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        Server resource = manager.servers()
            .getByResourceGroupWithResponse("testrg", "mysqltestserver", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withStorage(new Storage().withStorageSizeGB(30).withIops(200).withAutoGrow(EnableStatusEnum.DISABLED)
                .withAutoIoScaling(EnableStatusEnum.DISABLED))
            .withNetwork(new Network().withPublicNetworkAccess(EnableStatusEnum.DISABLED)).apply();
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
