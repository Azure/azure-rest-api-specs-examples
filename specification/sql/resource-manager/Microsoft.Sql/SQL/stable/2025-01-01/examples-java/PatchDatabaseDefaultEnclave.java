
import com.azure.resourcemanager.sql.models.AlwaysEncryptedEnclaveType;
import com.azure.resourcemanager.sql.models.DatabaseUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/PatchDatabaseDefaultEnclave.json
     */
    /**
     * Sample code: Updates a database with Default enclave type.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updatesADatabaseWithDefaultEnclaveType(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().update("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new DatabaseUpdate().withPreferredEnclaveType(AlwaysEncryptedEnclaveType.DEFAULT),
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
