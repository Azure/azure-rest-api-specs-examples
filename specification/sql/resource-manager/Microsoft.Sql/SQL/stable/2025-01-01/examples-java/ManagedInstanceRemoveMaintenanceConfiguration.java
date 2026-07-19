
import com.azure.resourcemanager.sql.models.ManagedInstanceUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceRemoveMaintenanceConfiguration.json
     */
    /**
     * Sample code: Remove maintenance policy from managed instance (select default maintenance policy).
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void removeMaintenancePolicyFromManagedInstanceSelectDefaultMaintenancePolicy(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().update("testrg", "testinstance",
            new ManagedInstanceUpdate().withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
            com.azure.core.util.Context.NONE);
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
