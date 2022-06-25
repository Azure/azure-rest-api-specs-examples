import com.azure.core.util.Context;
import com.azure.resourcemanager.digitaltwins.models.DigitalTwinsDescription;
import com.azure.resourcemanager.digitaltwins.models.DigitalTwinsIdentity;
import com.azure.resourcemanager.digitaltwins.models.DigitalTwinsIdentityType;
import java.util.HashMap;
import java.util.Map;

/** Samples for DigitalTwins Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsPatch_WithIdentity_example.json
     */
    /**
     * Sample code: Patch a DigitalTwinsInstance resource with identity.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void patchADigitalTwinsInstanceResourceWithIdentity(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        DigitalTwinsDescription resource =
            manager
                .digitalTwins()
                .getByResourceGroupWithResponse("resRg", "myDigitalTwinsService", Context.NONE)
                .getValue();
        resource.update().withIdentity(new DigitalTwinsIdentity().withType(DigitalTwinsIdentityType.NONE)).apply();
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
