
import com.azure.resourcemanager.devtestlabs.models.VirtualNetwork;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualNetworks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_Update
     * .json
     */
    /**
     * Sample code: VirtualNetworks_Update.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualNetworksUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        VirtualNetwork resource = manager.virtualNetworks().getWithResponse("resourceGroupName", "{labName}",
            "{virtualNetworkName}", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tagName1", "tagValue1")).apply();
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
