import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceSkuProperties;
import com.azure.resourcemanager.apimanagement.models.SkuType;
import com.azure.resourcemanager.apimanagement.models.VirtualNetworkConfiguration;
import com.azure.resourcemanager.apimanagement.models.VirtualNetworkType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApiManagementService CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateServiceInVnetWithPublicIP.json
     */
    /**
     * Sample code: ApiManagementCreateServiceInVnetWithPublicIP.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateServiceInVnetWithPublicIP(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .define("apimService1")
            .withRegion("East US 2 EUAP")
            .withExistingResourceGroup("rg1")
            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.PREMIUM).withCapacity(2))
            .withPublisherEmail("apim@autorestsdk.com")
            .withPublisherName("autorestsdk")
            .withTags(mapOf("tag1", "value1", "tag2", "value2", "tag3", "value3"))
            .withZones(Arrays.asList("1", "2"))
            .withPublicIpAddressId(
                "/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.Network/publicIPAddresses/apimazvnet")
            .withVirtualNetworkConfiguration(
                new VirtualNetworkConfiguration()
                    .withSubnetResourceId(
                        "/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.Network/virtualNetworks/apimcus/subnets/tenant"))
            .withVirtualNetworkType(VirtualNetworkType.EXTERNAL)
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
