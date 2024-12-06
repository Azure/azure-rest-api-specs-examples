
import com.azure.resourcemanager.digitaltwins.models.DigitalTwinsDescription;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DigitalTwins Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/
     * DigitalTwinsPatch_example.json
     */
    /**
     * Sample code: Patch a DigitalTwinsInstance resource.
     * 
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void
        patchADigitalTwinsInstanceResource(com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        DigitalTwinsDescription resource = manager.digitalTwins()
            .getByResourceGroupWithResponse("resRg", "myDigitalTwinsService", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("purpose", "dev")).apply();
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
