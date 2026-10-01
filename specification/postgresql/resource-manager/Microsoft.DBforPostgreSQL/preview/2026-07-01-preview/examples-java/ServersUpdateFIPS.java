
import com.azure.resourcemanager.postgresqlflexibleserver.models.FipsMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/ServersUpdateFIPS.json
     */
    /**
     * Sample code: Enable or disable FIPS mode on an existing server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void enableOrDisableFIPSModeOnAnExistingServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource = manager.servers()
            .getByResourceGroupWithResponse("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withFipsMode(FipsMode.DISABLED).apply();
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
