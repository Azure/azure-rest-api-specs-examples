
import com.azure.resourcemanager.storage.models.AccountImmutabilityPolicyProperties;
import com.azure.resourcemanager.storage.models.AccountImmutabilityPolicyState;
import com.azure.resourcemanager.storage.models.ImmutableStorageAccount;
import com.azure.resourcemanager.storage.models.StorageAccountUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountUpdateWithImmutabilityPolicy.json
     */
    /**
     * Sample code: StorageAccountUpdateWithImmutabilityPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountUpdateWithImmutabilityPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts()
            .updateWithResponse("res9407", "sto8596",
                new StorageAccountUpdateParameters()
                    .withImmutableStorageWithVersioning(new ImmutableStorageAccount().withEnabled(true)
                        .withImmutabilityPolicy(new AccountImmutabilityPolicyProperties()
                            .withImmutabilityPeriodSinceCreationInDays(15)
                            .withState(AccountImmutabilityPolicyState.LOCKED).withAllowProtectedAppendWrites(true))),
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
