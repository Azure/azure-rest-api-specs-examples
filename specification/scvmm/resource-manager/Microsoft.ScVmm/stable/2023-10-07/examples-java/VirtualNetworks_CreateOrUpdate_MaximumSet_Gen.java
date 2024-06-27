
import com.azure.resourcemanager.scvmm.models.ExtendedLocation;
import com.azure.resourcemanager.scvmm.models.VirtualNetworkProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualNetworks_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualNetworks_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualNetworksCreateOrUpdateMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().define("_").withRegion("fky").withExistingResourceGroup("rgscvmm")
            .withExtendedLocation(new ExtendedLocation().withType("customLocation").withName(
                "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"))
            .withTags(mapOf("key705", "fakeTokenPlaceholder"))
            .withProperties(new VirtualNetworkProperties().withInventoryItemId("bxn")
                .withUuid("12345678-1234-1234-1234-12345678abcd").withVmmServerId(
                    "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"))
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
