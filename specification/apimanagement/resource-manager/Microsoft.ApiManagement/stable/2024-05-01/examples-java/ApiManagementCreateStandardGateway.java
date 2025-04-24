
import com.azure.resourcemanager.apimanagement.models.ApiGatewaySkuType;
import com.azure.resourcemanager.apimanagement.models.ApiManagementGatewaySkuProperties;
import com.azure.resourcemanager.apimanagement.models.BackendConfiguration;
import com.azure.resourcemanager.apimanagement.models.BackendSubnetConfiguration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ApiGateway CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateStandardGateway.json
     */
    /**
     * Sample code: ApiManagementCreateStandardGateway.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateStandardGateway(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiGateways().define("apimGateway1").withRegion("South Central US").withExistingResourceGroup("rg1")
            .withSku(new ApiManagementGatewaySkuProperties().withName(ApiGatewaySkuType.STANDARD).withCapacity(1))
            .withTags(mapOf("Name", "Contoso", "Test", "User"))
            .withBackend(new BackendConfiguration().withSubnet(new BackendSubnetConfiguration().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vn1/subnets/sn1")))
            .create();
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
