
import com.azure.resourcemanager.oracledatabase.models.ExascaleDbStorageVault;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ExascaleDbStorageVaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/ExascaleDbStorageVaults_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbStorageVaults_Update_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void exascaleDbStorageVaultsUpdateMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        ExascaleDbStorageVault resource = manager.exascaleDbStorageVaults()
            .getByResourceGroupWithResponse("rgopenapi", "vmClusterName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key6179", "fakeTokenPlaceholder")).apply();
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
