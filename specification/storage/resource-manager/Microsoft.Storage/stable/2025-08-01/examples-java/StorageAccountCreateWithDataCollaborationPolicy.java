
import com.azure.resourcemanager.storage.models.Kind;
import com.azure.resourcemanager.storage.models.Sku;
import com.azure.resourcemanager.storage.models.SkuName;
import com.azure.resourcemanager.storage.models.StorageAccountCreateParameters;
import com.azure.resourcemanager.storage.models.StorageDataCollaborationPolicyProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountCreateWithDataCollaborationPolicy.json
     */
    /**
     * Sample code: StorageAccountCreateWithDataCollaborationPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountCreateWithDataCollaborationPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().create("res9101", "sto4445",
            new StorageAccountCreateParameters().withSku(new Sku().withName(SkuName.STANDARD_GRS))
                .withKind(Kind.STORAGE).withLocation("eastus").withDataCollaborationPolicyProperties(
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
