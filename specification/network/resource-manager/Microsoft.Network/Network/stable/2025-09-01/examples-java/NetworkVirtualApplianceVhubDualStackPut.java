
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.NetworkVirtualApplianceInner;
import com.azure.resourcemanager.network.models.IpVersion;
import com.azure.resourcemanager.network.models.ManagedServiceIdentity;
import com.azure.resourcemanager.network.models.ManagedServiceIdentityUserAssignedIdentities;
import com.azure.resourcemanager.network.models.ResourceIdentityType;
import com.azure.resourcemanager.network.models.VirtualApplianceSkuProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NetworkVirtualAppliances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkVirtualApplianceVhubDualStackPut.json
     */
    /**
     * Sample code: Create NVA in Virtual Hub for IPv4 and IPv6.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createNVAInVirtualHubForIPv4AndIPv6(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().createOrUpdate("rg1", "nva",
            new NetworkVirtualApplianceInner().withLocation("West US").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withIdentity(new ManagedServiceIdentity().withType(ResourceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(mapOf(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1",
                        new ManagedServiceIdentityUserAssignedIdentities())))
                .withNvaSku(new VirtualApplianceSkuProperties().withVendor("Cisco SDWAN").withBundledScaleUnit("1")
                    .withMarketPlaceVersion("12.1"))
                .withBootStrapConfigurationBlobs(Arrays
                    .asList("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"))
                .withVirtualHub(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"))
                .withCloudInitConfigurationBlobs(Arrays
                    .asList("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"))
                .withVirtualApplianceAsn(10000L).withAddressFamily(Arrays.asList(IpVersion.IPV4, IpVersion.IPV6)),
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
