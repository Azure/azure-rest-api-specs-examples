
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.NetworkVirtualApplianceInner;
import com.azure.resourcemanager.network.models.IpVersion;
import com.azure.resourcemanager.network.models.NetworkVirtualApplianceLoadBalancerConfiguration;
import com.azure.resourcemanager.network.models.NetworkVirtualApplianceLoadDistribution;
import com.azure.resourcemanager.network.models.VirtualApplianceSkuProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NetworkVirtualAppliances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/NetworkVirtualApplianceVhubSourceIPProtocolPut.json
     */
    /**
     * Sample code: Create or update a managed NVA in a virtual hub with SourceIPProtocol load distribution.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createOrUpdateAManagedNVAInAVirtualHubWithSourceIPProtocolLoadDistribution(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().createOrUpdate("rg1", "nva-vhub",
            new NetworkVirtualApplianceInner().withLocation("West US")
                .withNvaSku(new VirtualApplianceSkuProperties().withVendor("Cisco SDWAN").withBundledScaleUnit("1")
                    .withMarketPlaceVersion("12.1"))
                .withLoadBalancerConfiguration(new NetworkVirtualApplianceLoadBalancerConfiguration()
                    .withLoadDistribution(NetworkVirtualApplianceLoadDistribution.SOURCE_IPPROTOCOL)
                    .withIdleTimeoutInMinutes(15))
                .withVirtualHub(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"))
                .withVirtualApplianceAsn(10000L).withAddressFamily(Arrays.asList(IpVersion.IPV4)),
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
