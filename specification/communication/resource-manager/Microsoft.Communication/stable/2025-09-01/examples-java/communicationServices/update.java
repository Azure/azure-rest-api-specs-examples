
import com.azure.resourcemanager.communication.models.CommunicationServiceResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CommunicationServices Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/
     * communicationServices/update.json
     */
    /**
     * Sample code: Update resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void updateResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        CommunicationServiceResource resource
            = manager.communicationServices().getByResourceGroupWithResponse("MyResourceGroup",
                "MyCommunicationResource", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("newTag", "newVal")).apply();
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
