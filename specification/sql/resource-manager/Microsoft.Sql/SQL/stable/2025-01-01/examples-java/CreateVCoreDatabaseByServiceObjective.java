
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateVCoreDatabaseByServiceObjective.json
     */
    /**
     * Sample code: Creates a VCore database by specifying service objective name.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createsAVCoreDatabaseBySpecifyingServiceObjectiveName(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases()
            .createOrUpdate(
                "Default-SQL-SouthEastAsia", "testsvr", "testdb", new DatabaseInner().withLocation("southeastasia")
                    .withSku(new Sku().withName("BC").withFamily("Gen4").withCapacity(2)),
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
