
import com.azure.resourcemanager.compute.fluent.models.GalleryInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Galleries CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/
     * Gallery_Create.json
     */
    /**
     * Sample code: Create or update a simple gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().createOrUpdate("myResourceGroup",
            "myGalleryName",
            new GalleryInner().withLocation("West US").withDescription("This is the gallery description."),
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
