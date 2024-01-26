
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ManagedInstanceUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for ManagedInstances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceRemoveMaintenanceConfiguration.json
     */
    /**
     * Sample code: Remove maintenance policy from managed instance (select default maintenance policy).
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void removeMaintenancePolicyFromManagedInstanceSelectDefaultMaintenancePolicy(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().update("testrg", "testinstance",
            new ManagedInstanceUpdate().withMaintenanceConfigurationId(
                "/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
            Context.NONE);
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
