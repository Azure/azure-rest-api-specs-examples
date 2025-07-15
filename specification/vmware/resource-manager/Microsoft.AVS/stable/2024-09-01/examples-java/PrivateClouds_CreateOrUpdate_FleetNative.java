
import com.azure.resourcemanager.avs.models.DnsZoneType;
import com.azure.resourcemanager.avs.models.ManagementCluster;
import com.azure.resourcemanager.avs.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PrivateClouds CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateClouds_CreateOrUpdate_FleetNative.json
     */
    /**
     * Sample code: PrivateClouds_CreateOrUpdate_FleetNative.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsCreateOrUpdateFleetNative(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().define("cloud1").withRegion("eastus2").withExistingResourceGroup("group1")
            .withSku(new Sku().withName("AV64")).withTags(mapOf())
            .withManagementCluster(new ManagementCluster().withClusterSize(4)).withNetworkBlock("192.168.48.0/22")
            .withVirtualNetworkId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/vnet")
            .withDnsZoneType(DnsZoneType.PRIVATE).create();
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
