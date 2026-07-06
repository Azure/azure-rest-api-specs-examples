
import com.azure.resourcemanager.storage.models.StorageAccountUpdateParameters;
import com.azure.resourcemanager.storage.models.StorageDataCollaborationPolicyProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountUpdateWithDataCollaborationPolicy.json
     */
    /**
     * Sample code: StorageAccountUpdateWithDataCollaborationPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountUpdateWithDataCollaborationPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().updateWithResponse("res9407", "sto8596",
            new StorageAccountUpdateParameters().withDataCollaborationPolicyProperties(
                new StorageDataCollaborationPolicyProperties().withAllowStorageConnectors(true)
                    .withAllowStorageDataShares(true).withAllowCrossTenantDataSharing(false)),
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
