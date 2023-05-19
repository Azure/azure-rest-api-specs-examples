import com.azure.resourcemanager.mysqlflexibleserver.models.EnableStatusEnum;
import com.azure.resourcemanager.mysqlflexibleserver.models.Server;
import com.azure.resourcemanager.mysqlflexibleserver.models.Storage;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2022-09-30-preview/examples/ServerUpdate.json
     */
    /**
     * Sample code: Update a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void updateAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        Server resource =
            manager
                .servers()
                .getByResourceGroupWithResponse("testrg", "mysqltestserver", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withStorage(
                new Storage()
                    .withStorageSizeGB(30)
                    .withIops(200)
                    .withAutoGrow(EnableStatusEnum.DISABLED)
                    .withAutoIoScaling(EnableStatusEnum.DISABLED))
            .apply();
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
