
import com.azure.resourcemanager.compute.models.DedicatedHostGroupUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DedicatedHostGroups Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * dedicatedHostExamples/DedicatedHostGroup_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_Update_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostGroupUpdateMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHostGroups().updateWithResponse("rgcompute",
            "aaaaaaaaaaa", new DedicatedHostGroupUpdate(), com.azure.core.util.Context.NONE);
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
