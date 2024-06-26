import java.util.HashMap;
import java.util.Map;

/** Samples for StorageMovers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/StorageMovers_CreateOrUpdate.json
     */
    /**
     * Sample code: StorageMovers_CreateOrUpdate.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void storageMoversCreateOrUpdate(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .storageMovers()
            .define("examples-storageMoverName")
            .withRegion("eastus2")
            .withExistingResourceGroup("examples-rg")
            .withTags(mapOf("key1", "value1", "key2", "value2"))
            .withDescription("Example Storage Mover Description")
            .create();
    }

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
