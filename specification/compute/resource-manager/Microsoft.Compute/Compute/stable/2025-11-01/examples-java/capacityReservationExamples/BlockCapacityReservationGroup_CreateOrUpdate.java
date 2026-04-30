
import com.azure.resourcemanager.compute.fluent.models.CapacityReservationGroupInner;
import com.azure.resourcemanager.compute.models.ReservationType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CapacityReservationGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/capacityReservationExamples/BlockCapacityReservationGroup_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a block capacity reservation group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createOrUpdateABlockCapacityReservationGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservationGroups().createOrUpdateWithResponse("myResourceGroup",
            "blockCapacityReservationGroup",
            new CapacityReservationGroupInner().withLocation("westus").withTags(mapOf("department", "finance"))
                .withZones(Arrays.asList("1", "2")).withReservationType(ReservationType.BLOCK),
            com.azure.core.util.Context.NONE);
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
