import com.azure.resourcemanager.azurestackhci.models.StorageContainers;
import java.util.HashMap;
import java.util.Map;

/** Samples for StorageContainersOperation Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateStorageContainer.json
     */
    /**
     * Sample code: UpdateStorageContainer.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void updateStorageContainer(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        StorageContainers resource =
            manager
                .storageContainersOperations()
                .getByResourceGroupWithResponse("test-rg", "Default_Container", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("additionalProperties", "sample")).apply();
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
