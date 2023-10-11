import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateModeForUpdate;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import com.azure.resourcemanager.postgresqlflexibleserver.models.ServerVersion;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/ServerUpdateWithMajorVersionUpgrade.json
     */
    /**
     * Sample code: ServerUpdateWithMajorVersionUpgrade.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverUpdateWithMajorVersionUpgrade(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource =
            manager
                .servers()
                .getByResourceGroupWithResponse("testrg", "pgtestsvc4", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withVersion(ServerVersion.ONE_FOUR).withCreateMode(CreateModeForUpdate.UPDATE).apply();
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
