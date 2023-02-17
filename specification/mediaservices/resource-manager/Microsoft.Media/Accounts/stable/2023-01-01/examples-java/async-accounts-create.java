import com.azure.resourcemanager.mediaservices.models.StorageAccount;
import com.azure.resourcemanager.mediaservices.models.StorageAccountType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Mediaservices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/async-accounts-create.json
     */
    /**
     * Sample code: Create a Media Services account.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createAMediaServicesAccount(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .mediaservices()
            .define("contososports")
            .withRegion("South Central US")
            .withExistingResourceGroup("contosorg")
            .withTags(mapOf("key1", "value1", "key2", "value2"))
            .withStorageAccounts(
                Arrays
                    .asList(
                        new StorageAccount()
                            .withId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosorg/providers/Microsoft.Storage/storageAccounts/teststorageaccount")
                            .withType(StorageAccountType.PRIMARY)))
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
