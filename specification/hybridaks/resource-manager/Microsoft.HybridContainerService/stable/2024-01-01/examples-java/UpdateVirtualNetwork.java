
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetwork;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualNetworks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * UpdateVirtualNetwork.json
     */
    /**
     * Sample code: UpdateVirtualNetwork.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        updateVirtualNetwork(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        VirtualNetwork resource = manager.virtualNetworks().getByResourceGroupWithResponse("test-arcappliance-resgrp",
            "test-vnet-static", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("additionalProperties", "sample")).apply();
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
