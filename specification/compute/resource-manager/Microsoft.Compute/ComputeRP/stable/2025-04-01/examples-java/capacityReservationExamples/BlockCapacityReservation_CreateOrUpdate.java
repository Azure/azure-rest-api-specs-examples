
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
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * capacityReservationExamples/BlockCapacityReservation_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a block capacity reservation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateABlockCapacityReservation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservations().createOrUpdate("myResourceGroup",
            "blockCapacityReservationGroup", "blockCapacityReservation",
            new CapacityReservationInner().withLocation("westus").withTags(mapOf("department", "HR"))
                .withSku(new Sku().withName("Standard_ND96isr_H100_v5").withCapacity(1L)).withZones(Arrays.asList("1"))
                .withScheduleProfile(new ScheduleProfile().withStart("2025-08-01").withEnd("2025-08-02")),
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
