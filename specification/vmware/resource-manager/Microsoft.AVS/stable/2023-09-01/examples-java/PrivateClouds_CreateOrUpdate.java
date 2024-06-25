
import com.azure.resourcemanager.avs.models.ManagementCluster;
import com.azure.resourcemanager.avs.models.PrivateCloudIdentity;
import com.azure.resourcemanager.avs.models.ResourceIdentityType;
import com.azure.resourcemanager.avs.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PrivateClouds CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/PrivateClouds_CreateOrUpdate.json
     */
    /**
     * Sample code: PrivateClouds_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().define("cloud1").withRegion("eastus2").withExistingResourceGroup("group1")
            .withSku(new Sku().withName("AV36")).withTags(mapOf())
            .withIdentity(new PrivateCloudIdentity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withManagementCluster(new ManagementCluster().withClusterSize(4)).withNetworkBlock("192.168.48.0/22")
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
