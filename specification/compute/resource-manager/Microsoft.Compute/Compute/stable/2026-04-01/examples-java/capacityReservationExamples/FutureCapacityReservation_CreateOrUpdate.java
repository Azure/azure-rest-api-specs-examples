
import com.azure.resourcemanager.compute.fluent.models.CapacityReservationInner;
import com.azure.resourcemanager.compute.models.ScheduleProfile;
import com.azure.resourcemanager.compute.models.Sku;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CapacityReservations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/capacityReservationExamples/FutureCapacityReservation_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a Future capacity reservation.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createOrUpdateAFutureCapacityReservation(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().createOrUpdate("myResourceGroup",
            "futureCapacityReservationGroup", "futureCapacityReservation",
            new CapacityReservationInner().withLocation("westus").withTags(mapOf("department", "HR"))
                .withSku(new Sku().withName("Standard_DS1_v2").withCapacity(4L)).withZones(Arrays.asList("1"))
                .withScheduleProfile(
                    new ScheduleProfile().withStart("2026-08-01T12:00:00Z").withMinimumCommitmentDays(30)),
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
