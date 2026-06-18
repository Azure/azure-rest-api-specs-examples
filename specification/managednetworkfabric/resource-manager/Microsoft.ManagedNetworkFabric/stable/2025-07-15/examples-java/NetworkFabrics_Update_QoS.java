
import com.azure.resourcemanager.managednetworkfabric.models.NetworkFabric;
import com.azure.resourcemanager.managednetworkfabric.models.QosConfigurationState;
import com.azure.resourcemanager.managednetworkfabric.models.QosPatchProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NetworkFabrics Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_Update_QoS.json
     */
    /**
     * Sample code: NetworkFabrics_Update_QoS_Enable.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsUpdateQoSEnable(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        NetworkFabric resource = manager.networkFabrics()
            .getByResourceGroupWithResponse("example-rg", "example-fabric", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withQosConfiguration(new QosPatchProperties().withQosConfigurationState(QosConfigurationState.ENABLED))
            .apply();
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
