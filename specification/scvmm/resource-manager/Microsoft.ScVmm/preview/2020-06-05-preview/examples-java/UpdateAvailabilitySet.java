import com.azure.core.util.Context;
import com.azure.resourcemanager.scvmm.models.AvailabilitySet;
import java.util.HashMap;
import java.util.Map;

/** Samples for AvailabilitySets Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/UpdateAvailabilitySet.json
     */
    /**
     * Sample code: UpdateAvailabilitySet.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void updateAvailabilitySet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        AvailabilitySet resource =
            manager
                .availabilitySets()
                .getByResourceGroupWithResponse("testrg", "HRAvailabilitySet", Context.NONE)
                .getValue();
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
