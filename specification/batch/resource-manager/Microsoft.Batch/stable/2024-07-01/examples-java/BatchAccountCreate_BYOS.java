
import com.azure.resourcemanager.batch.models.AutoStorageBaseProperties;
import com.azure.resourcemanager.batch.models.KeyVaultReference;
import com.azure.resourcemanager.batch.models.PoolAllocationMode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BatchAccount Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/BatchAccountCreate_BYOS.json
     */
    /**
     * Sample code: BatchAccountCreate_BYOS.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountCreateBYOS(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().define("sampleacct").withRegion("japaneast")
            .withExistingResourceGroup("default-azurebatch-japaneast")
            .withAutoStorage(new AutoStorageBaseProperties().withStorageAccountId(
                "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"))
            .withPoolAllocationMode(PoolAllocationMode.USER_SUBSCRIPTION)
            .withKeyVaultReference(new KeyVaultReference().withId(
                "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.KeyVault/vaults/sample")
                .withUrl("http://sample.vault.azure.net/"))
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
