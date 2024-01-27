
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerAutomaticTuningInner;
import com.azure.resourcemanager.sql.models.AutomaticTuningServerMode;
import java.util.HashMap;
import java.util.Map;

/** Samples for ServerAutomaticTuning Update. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerAutomaticTuningUpdateMin.json
     */
    /**
     * Sample code: Updates server automatic tuning settings with minimal properties.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatesServerAutomaticTuningSettingsWithMinimalProperties(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAutomaticTunings().updateWithResponse(
            "default-sql-onebox", "testsvr11",
            new ServerAutomaticTuningInner().withDesiredState(AutomaticTuningServerMode.AUTO), Context.NONE);
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
