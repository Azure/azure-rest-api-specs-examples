
import com.azure.resourcemanager.scvmm.models.VmmServer;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VmmServers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmmServers_Update_MaximumSet_Gen.
     * json
     */
    /**
     * Sample code: VmmServers_Update_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void vmmServersUpdateMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        VmmServer resource = manager.vmmServers()
            .getByResourceGroupWithResponse("rgscvmm", "Y", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key7187", "fakeTokenPlaceholder")).apply();
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
