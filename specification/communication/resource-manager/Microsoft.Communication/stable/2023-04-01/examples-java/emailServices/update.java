
import com.azure.resourcemanager.communication.models.EmailServiceResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for EmailServices Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2023-04-01/examples/emailServices/
     * update.json
     */
    /**
     * Sample code: Update EmailService resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        updateEmailServiceResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        EmailServiceResource resource = manager.emailServices().getByResourceGroupWithResponse("MyResourceGroup",
            "MyEmailServiceResource", com.azure.core.util.Context.NONE).getValue();
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
