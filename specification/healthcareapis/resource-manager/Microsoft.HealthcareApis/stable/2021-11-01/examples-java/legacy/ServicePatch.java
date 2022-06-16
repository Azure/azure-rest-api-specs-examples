import com.azure.core.util.Context;
import com.azure.resourcemanager.healthcareapis.models.ServicesDescription;
import java.util.HashMap;
import java.util.Map;

/** Samples for Services Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServicePatch.json
     */
    /**
     * Sample code: Patch service.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void patchService(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        ServicesDescription resource =
            manager.services().getByResourceGroupWithResponse("rg1", "service1", Context.NONE).getValue();
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
