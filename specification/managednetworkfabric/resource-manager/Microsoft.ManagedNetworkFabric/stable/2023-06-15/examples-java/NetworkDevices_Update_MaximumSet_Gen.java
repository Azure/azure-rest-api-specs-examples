
import com.azure.resourcemanager.managednetworkfabric.models.NetworkDevice;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NetworkDevices Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * NetworkDevices_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkDevices_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesUpdateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        NetworkDevice resource = manager.networkDevices()
            .getByResourceGroupWithResponse("example-rg", "example-device", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("keyID", "fakeTokenPlaceholder")).withHostname("NFA-Device")
            .withSerialNumber("Vendor;DCS-7280XXX-24;12.05;JPE2111XXXX").withAnnotation("annotation").apply();
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
