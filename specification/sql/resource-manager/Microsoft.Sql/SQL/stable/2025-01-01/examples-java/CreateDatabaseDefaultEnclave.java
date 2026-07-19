
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.AlwaysEncryptedEnclaveType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateDatabaseDefaultEnclave.json
     */
    /**
     * Sample code: Creates a database with Default enclave type.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createsADatabaseWithDefaultEnclaveType(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases()
            .createOrUpdate("Default-SQL-SouthEastAsia", "testsvr", "testdb", new DatabaseInner()
                .withLocation("southeastasia").withPreferredEnclaveType(AlwaysEncryptedEnclaveType.DEFAULT),
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
