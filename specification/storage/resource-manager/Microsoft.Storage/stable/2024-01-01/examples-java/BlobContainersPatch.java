
import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;
import com.azure.resourcemanager.storage.models.PublicAccess;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BlobContainers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/BlobContainersPatch.json
     */
    /**
     * Sample code: UpdateContainers.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateContainers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobContainers()
            .updateWithResponse("res3376", "sto328", "container6185", new BlobContainerInner()
                .withPublicAccess(PublicAccess.CONTAINER).withMetadata(mapOf("metadata", "true")),
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
