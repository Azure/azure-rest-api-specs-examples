
import com.azure.resourcemanager.managednetworkfabric.models.NetworkPacketBroker;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NetworkPacketBrokers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * NetworkPacketBrokers_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkPacketBrokers_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkPacketBrokersUpdateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        NetworkPacketBroker resource = manager.networkPacketBrokers().getByResourceGroupWithResponse("example-rg",
            "example-networkPacketBroker", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key8772", "fakeTokenPlaceholder")).apply();
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
