
import com.azure.resourcemanager.compute.models.CapacityReservationUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CapacityReservations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * capacityReservationExamples/CapacityReservation_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservation_Update_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void capacityReservationUpdateMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservations().update("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaa", "aaa", new CapacityReservationUpdate(), com.azure.core.util.Context.NONE);
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
