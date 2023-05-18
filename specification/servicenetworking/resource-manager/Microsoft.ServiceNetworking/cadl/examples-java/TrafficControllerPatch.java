import com.azure.resourcemanager.servicenetworking.models.TrafficController;
import java.util.HashMap;
import java.util.Map;

/** Samples for TrafficControllerInterface Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/TrafficControllerPatch.json
     */
    /**
     * Sample code: Patch Traffic Controller.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void patchTrafficController(
        com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        TrafficController resource =
            manager
                .trafficControllerInterfaces()
                .getByResourceGroupWithResponse("rg1", "tc1", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("key1", "value1")).apply();
    }

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
