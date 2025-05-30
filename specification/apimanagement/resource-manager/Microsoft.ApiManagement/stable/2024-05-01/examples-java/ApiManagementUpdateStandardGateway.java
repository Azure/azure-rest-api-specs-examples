
import com.azure.resourcemanager.apimanagement.models.ApiGatewaySkuType;
import com.azure.resourcemanager.apimanagement.models.ApiManagementGatewayResource;
import com.azure.resourcemanager.apimanagement.models.ApiManagementGatewaySkuPropertiesForPatch;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ApiGateway Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateStandardGateway.json
     */
    /**
     * Sample code: ApiManagementUpdateStandardGateway.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateStandardGateway(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiManagementGatewayResource resource = manager.apiGateways()
            .getByResourceGroupWithResponse("rg1", "apimGateway1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("Name", "Contoso", "Test", "User"))
            .withSku(
                new ApiManagementGatewaySkuPropertiesForPatch().withName(ApiGatewaySkuType.STANDARD).withCapacity(10))
            .apply();
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
