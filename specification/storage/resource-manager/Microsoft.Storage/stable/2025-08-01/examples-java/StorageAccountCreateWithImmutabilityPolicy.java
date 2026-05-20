
import com.azure.resourcemanager.storage.models.AccountImmutabilityPolicyProperties;
import com.azure.resourcemanager.storage.models.AccountImmutabilityPolicyState;
import com.azure.resourcemanager.storage.models.ExtendedLocation;
import com.azure.resourcemanager.storage.models.ExtendedLocationTypes;
import com.azure.resourcemanager.storage.models.ImmutableStorageAccount;
import com.azure.resourcemanager.storage.models.Kind;
import com.azure.resourcemanager.storage.models.Sku;
import com.azure.resourcemanager.storage.models.SkuName;
import com.azure.resourcemanager.storage.models.StorageAccountCreateParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountCreateWithImmutabilityPolicy.json
     */
    /**
     * Sample code: StorageAccountCreateWithImmutabilityPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountCreateWithImmutabilityPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts()
            .create("res9101", "sto4445",
                new StorageAccountCreateParameters().withSku(new Sku().withName(SkuName.STANDARD_GRS))
                    .withKind(Kind.STORAGE).withLocation("eastus")
                    .withExtendedLocation(
                        new ExtendedLocation().withName("losangeles001").withType(ExtendedLocationTypes.EDGE_ZONE))
                    .withImmutableStorageWithVersioning(new ImmutableStorageAccount().withEnabled(true)
                        .withImmutabilityPolicy(new AccountImmutabilityPolicyProperties()
                            .withImmutabilityPeriodSinceCreationInDays(15)
                            .withState(AccountImmutabilityPolicyState.UNLOCKED).withAllowProtectedAppendWrites(true))),
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
