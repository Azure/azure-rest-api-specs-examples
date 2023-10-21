import com.azure.resourcemanager.azurestackhci.models.GalleryImages;
import java.util.HashMap;
import java.util.Map;

/** Samples for GalleryImagesOperation Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateGalleryImage.json
     */
    /**
     * Sample code: UpdateGalleryImage.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void updateGalleryImage(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        GalleryImages resource =
            manager
                .galleryImagesOperations()
                .getByResourceGroupWithResponse("test-rg", "test-gallery-image", com.azure.core.util.Context.NONE)
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
