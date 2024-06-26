
import com.azure.resourcemanager.scvmm.models.Cloud;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clouds Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/Clouds_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: Clouds_Update_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void cloudsUpdateMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        Cloud resource = manager.clouds()
            .getByResourceGroupWithResponse("rgscvmm", "_", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
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
