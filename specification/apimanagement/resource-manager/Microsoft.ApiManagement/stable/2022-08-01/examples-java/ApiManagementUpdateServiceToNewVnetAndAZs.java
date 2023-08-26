import com.azure.resourcemanager.apimanagement.models.AdditionalLocation;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceResource;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceSkuProperties;
import com.azure.resourcemanager.apimanagement.models.SkuType;
import com.azure.resourcemanager.apimanagement.models.VirtualNetworkConfiguration;
import com.azure.resourcemanager.apimanagement.models.VirtualNetworkType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApiManagementService Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateServiceToNewVnetAndAZs.json
     */
    /**
     * Sample code: ApiManagementUpdateServiceToNewVnetAndAvailabilityZones.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateServiceToNewVnetAndAvailabilityZones(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiManagementServiceResource resource =
            manager
                .apiManagementServices()
                .getByResourceGroupWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.PREMIUM).withCapacity(3))
            .withZones(Arrays.asList("1", "2", "3"))
            .withPublicIpAddressId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/publicip-apim-japan-east")
            .withVirtualNetworkConfiguration(
                new VirtualNetworkConfiguration()
                    .withSubnetResourceId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-apim-japaneast/subnets/apim2"))
            .withAdditionalLocations(
                Arrays
                    .asList(
                        new AdditionalLocation()
                            .withLocation("Australia East")
                            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.PREMIUM).withCapacity(3))
                            .withZones(Arrays.asList("1", "2", "3"))
                            .withPublicIpAddressId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/apim-australia-east-publicip")
                            .withVirtualNetworkConfiguration(
                                new VirtualNetworkConfiguration()
                                    .withSubnetResourceId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/apimaeavnet/subnets/default"))))
            .withVirtualNetworkType(VirtualNetworkType.EXTERNAL)
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
