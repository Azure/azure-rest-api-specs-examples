
import com.azure.resourcemanager.networkcloud.models.AdministrativeCredentials;
import com.azure.resourcemanager.networkcloud.models.ExtendedLocation;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAppliances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/
     * StorageAppliances_Create.json
     */
    /**
     * Sample code: Create or update storage appliance.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        createOrUpdateStorageAppliance(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.storageAppliances().define("storageApplianceName").withRegion("location")
            .withExistingResourceGroup("resourceGroupName")
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName")
                .withType("CustomLocation"))
            .withAdministratorCredentials(
                new AdministrativeCredentials().withPassword("fakeTokenPlaceholder").withUsername("adminUser"))
            .withRackId(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName")
            .withRackSlot(1L).withSerialNumber("BM1219XXX").withStorageApplianceSkuId("684E-3B16-399E")
            .withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder")).create();
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
