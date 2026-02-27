
import com.azure.resourcemanager.network.fluent.models.NetworkVirtualApplianceInner;
import com.azure.resourcemanager.network.models.ManagedServiceIdentity;
import com.azure.resourcemanager.network.models.ManagedServiceIdentityUserAssignedIdentities;
import com.azure.resourcemanager.network.models.NvaInVnetSubnetReferenceProperties;
import com.azure.resourcemanager.network.models.NvaInterfaceConfigurationsProperties;
import com.azure.resourcemanager.network.models.NvaNicType;
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
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkVirtualApplianceVnetBasicPut.json
     */
    /**
     * Sample code: Create NVA in VNet with PrivateNic &amp; PublicNic.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNVAInVNetWithPrivateNicPublicNic(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkVirtualAppliances().createOrUpdate("rg1", "nva",
            new NetworkVirtualApplianceInner().withLocation("West US").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withIdentity(new ManagedServiceIdentity().withType(ResourceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(mapOf(
                        "/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1",
                        new ManagedServiceIdentityUserAssignedIdentities())))
                .withNvaSku(new VirtualApplianceSkuProperties().withVendor("Cisco SDWAN").withBundledScaleUnit("1")
                    .withMarketPlaceVersion("latest"))
                .withBootStrapConfigurationBlobs(Arrays
                    .asList("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"))
                .withCloudInitConfigurationBlobs(Arrays
                    .asList("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"))
                .withVirtualApplianceAsn(10000L)
                .withNvaInterfaceConfigurations(Arrays.asList(new NvaInterfaceConfigurationsProperties()
                    .withSubnet(new NvaInVnetSubnetReferenceProperties().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"))
                    .withType(Arrays.asList(NvaNicType.PRIVATE_NIC)).withName("dataInterface"),
                    new NvaInterfaceConfigurationsProperties()
                        .withSubnet(new NvaInVnetSubnetReferenceProperties().withId(
                            "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet2"))
                        .withType(Arrays.asList(NvaNicType.PUBLIC_NIC)).withName("managementInterface"))),
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
