
import com.azure.resourcemanager.batch.models.AutoStorageBaseProperties;
import com.azure.resourcemanager.batch.models.BatchAccountIdentity;
import com.azure.resourcemanager.batch.models.ResourceIdentityType;
import com.azure.resourcemanager.batch.models.UserAssignedIdentities;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BatchAccount Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/BatchAccountCreate_UserAssignedIdentity.json
     */
    /**
     * Sample code: BatchAccountCreate_UserAssignedIdentity.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountCreateUserAssignedIdentity(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().define("sampleacct").withRegion("japaneast")
            .withExistingResourceGroup("default-azurebatch-japaneast")
            .withIdentity(new BatchAccountIdentity().withType(ResourceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                    new UserAssignedIdentities())))
            .withAutoStorage(new AutoStorageBaseProperties().withStorageAccountId(
                "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"))
            .create();
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
