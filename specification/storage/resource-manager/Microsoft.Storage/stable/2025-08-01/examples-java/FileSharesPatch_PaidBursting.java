
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;
import com.azure.resourcemanager.storage.models.FileSharePropertiesFileSharePaidBursting;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for FileShares Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileSharesPatch_PaidBursting.json
     */
    /**
     * Sample code: UpdateSharePaidBursting.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void updateSharePaidBursting(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().updateWithResponse("res3376", "sto328", "share6185",
            new FileShareInner()
                .withFileSharePaidBursting(new FileSharePropertiesFileSharePaidBursting().withPaidBurstingEnabled(true)
                    .withPaidBurstingMaxIops(102400).withPaidBurstingMaxBandwidthMibps(10340)),
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
