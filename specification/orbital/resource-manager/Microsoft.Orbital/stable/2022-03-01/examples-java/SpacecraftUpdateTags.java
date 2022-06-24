import com.azure.core.util.Context;
import com.azure.resourcemanager.orbital.models.Spacecraft;
import java.util.HashMap;
import java.util.Map;

/** Samples for Spacecrafts UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftUpdateTags.json
     */
    /**
     * Sample code: Update Spacecraft tags.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void updateSpacecraftTags(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        Spacecraft resource =
            manager.spacecrafts().getByResourceGroupWithResponse("contoso-Rgp", "CONTOSO_SAT", Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
