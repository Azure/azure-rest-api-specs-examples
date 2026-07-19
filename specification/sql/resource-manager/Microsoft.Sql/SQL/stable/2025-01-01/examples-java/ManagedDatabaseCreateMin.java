
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedDatabases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseCreateMin.json
     */
    /**
     * Sample code: Creates a new managed database with minimal properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createsANewManagedDatabaseWithMinimalProperties(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().createOrUpdate("Default-SQL-SouthEastAsia", "managedInstance",
            "managedDatabase", new ManagedDatabaseInner().withLocation("southeastasia"),
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
