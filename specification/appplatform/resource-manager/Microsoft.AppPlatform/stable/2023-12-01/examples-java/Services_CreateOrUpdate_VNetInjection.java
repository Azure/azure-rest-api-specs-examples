
import com.azure.resourcemanager.appplatform.fluent.models.ServiceResourceInner;
import com.azure.resourcemanager.appplatform.models.ClusterResourceProperties;
import com.azure.resourcemanager.appplatform.models.IngressConfig;
import com.azure.resourcemanager.appplatform.models.NetworkProfile;
import com.azure.resourcemanager.appplatform.models.ServiceVNetAddons;
import com.azure.resourcemanager.appplatform.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Services_CreateOrUpdate_VNetInjection.json
     */
    /**
     * Sample code: Services_CreateOrUpdate_VNetInjection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesCreateOrUpdateVNetInjection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getServices().createOrUpdate("myResourceGroup", "myservice",
            new ServiceResourceInner().withLocation("eastus").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withProperties(new ClusterResourceProperties().withNetworkProfile(new NetworkProfile()
                    .withServiceRuntimeSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/serviceRuntime")
                    .withAppSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/apps")
                    .withServiceCidr("10.8.0.0/16,10.244.0.0/16,10.245.0.1/16")
                    .withServiceRuntimeNetworkResourceGroup("my-service-runtime-network-rg")
                    .withAppNetworkResourceGroup("my-app-network-rg")
                    .withIngressConfig(new IngressConfig().withReadTimeoutInSeconds(300))).withVnetAddons(
                        new ServiceVNetAddons().withLogStreamPublicEndpoint(true).withDataPlanePublicEndpoint(true)))
                .withSku(new Sku().withName("S0").withTier("Standard")),
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
