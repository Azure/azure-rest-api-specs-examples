
import com.azure.resourcemanager.containerservicefleet.models.Fleet;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Fleets Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/
     * examples/Fleets_PatchTags.json
     */
    /**
     * Sample code: Update a Fleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        updateAFleet(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        Fleet resource = manager.fleets()
            .getByResourceGroupWithResponse("rg1", "fleet1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("env", "prod", "tier", "secure")).withIfMatch("dfjkwelr7384").apply();
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
