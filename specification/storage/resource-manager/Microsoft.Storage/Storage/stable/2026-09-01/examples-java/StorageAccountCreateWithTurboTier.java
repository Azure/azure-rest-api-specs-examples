
import com.azure.resourcemanager.storage.models.AccessTier;
import com.azure.resourcemanager.storage.models.Kind;
import com.azure.resourcemanager.storage.models.Sku;
import com.azure.resourcemanager.storage.models.SkuName;
import com.azure.resourcemanager.storage.models.StorageAccountCreateParameters;
import com.azure.resourcemanager.storage.models.TurboTier;
import com.azure.resourcemanager.storage.models.TurboTierStatus;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/StorageAccountCreateWithTurboTier.json
     */
    /**
     * Sample code: StorageAccountCreateWithTurboTier.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountCreateWithTurboTier(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().create("res9101", "staturbocontoso01",
            new StorageAccountCreateParameters().withSku(new Sku().withName(SkuName.STANDARD_LRS))
                .withKind(Kind.STORAGE_V2).withLocation("eastus2").withAccessTier(AccessTier.HOT)
                .withTurboTier(new TurboTier().withStatus(TurboTierStatus.ENABLED).withTargetPercent(10)),
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
