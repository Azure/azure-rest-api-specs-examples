
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ApiManagementService Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateServiceDisableTls10.json
     */
    /**
     * Sample code: ApiManagementUpdateServiceDisableTls10.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateServiceDisableTls10(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiManagementServiceResource resource = manager.apiManagementServices()
            .getByResourceGroupWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withCustomProperties(
            mapOf("Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10", "false")).apply();
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
