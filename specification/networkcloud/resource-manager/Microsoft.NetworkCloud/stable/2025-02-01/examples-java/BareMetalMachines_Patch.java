
import com.azure.resourcemanager.networkcloud.models.BareMetalMachine;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BareMetalMachines Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/
     * BareMetalMachines_Patch.json
     */
    /**
     * Sample code: Patch bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        BareMetalMachine resource = manager.bareMetalMachines().getByResourceGroupWithResponse("resourceGroupName",
            "bareMetalMachineName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withMachineDetails("machinedetails").apply();
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
