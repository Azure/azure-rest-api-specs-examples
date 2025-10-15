
import com.azure.resourcemanager.oracledatabase.models.AutonomousDatabase;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AutonomousDatabases Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabase_patch.json
     */
    /**
     * Sample code: AutonomousDatabases_update.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesUpdate(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        AutonomousDatabase resource = manager.autonomousDatabases()
            .getByResourceGroupWithResponse("rg000", "databasedb1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
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
