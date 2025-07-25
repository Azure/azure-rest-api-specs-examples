
import com.azure.resourcemanager.networkcloud.models.AdministrativeCredentials;
import com.azure.resourcemanager.networkcloud.models.ExtendedLocation;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BareMetalMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/
     * BareMetalMachines_Create.json
     */
    /**
     * Sample code: Create or update bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        createOrUpdateBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().define("bareMetalMachineName").withRegion("location")
            .withExistingResourceGroup("resourceGroupName")
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName")
                .withType("CustomLocation"))
            .withBmcConnectionString("bmcconnectionstring")
            .withBmcCredentials(
                new AdministrativeCredentials().withPassword("fakeTokenPlaceholder").withUsername("bmcuser"))
            .withBmcMacAddress("00:00:4f:00:57:00").withBootMacAddress("00:00:4e:00:58:af")
            .withMachineDetails("User-provided machine details.").withMachineName("r01c001")
            .withMachineSkuId("684E-3B16-399E")
            .withRackId(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName")
            .withRackSlot(1L).withSerialNumber("BM1219XXX")
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
