
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.models.HyperVGenerationTypes;
import com.azure.resourcemanager.compute.models.ImageUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Images Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/imageExamples/
     * Image_Update.json
     */
    /**
     * Sample code: Updates tags of an Image.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatesTagsOfAnImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getImages().update("myResourceGroup", "myImage",
            new ImageUpdate().withTags(mapOf("department", "HR")).withSourceVirtualMachine(new SubResource().withId(
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"))
                .withHyperVGeneration(HyperVGenerationTypes.V1),
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
