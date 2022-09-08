import com.azure.core.util.Context;
import com.azure.resourcemanager.agrifood.models.FarmBeats;
import com.azure.resourcemanager.agrifood.models.FarmBeatsUpdateProperties;
import com.azure.resourcemanager.agrifood.models.Identity;
import com.azure.resourcemanager.agrifood.models.ResourceIdentityType;
import com.azure.resourcemanager.agrifood.models.SensorIntegration;
import java.util.HashMap;
import java.util.Map;

/** Samples for FarmBeatsModels Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_UpdateWithSensor.json
     */
    /**
     * Sample code: FarmBeatsModels_UpdateWithSensor.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void farmBeatsModelsUpdateWithSensor(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        FarmBeats resource =
            manager
                .farmBeatsModels()
                .getByResourceGroupWithResponse("examples-rg", "examples-farmBeatsResourceName", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("key1", "value1", "key2", "value2"))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withProperties(
                new FarmBeatsUpdateProperties().withSensorIntegration(new SensorIntegration().withEnabled("True")))
            .apply();
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
