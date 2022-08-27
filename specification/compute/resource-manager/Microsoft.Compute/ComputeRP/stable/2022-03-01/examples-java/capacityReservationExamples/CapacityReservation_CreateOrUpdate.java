import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.CapacityReservationInner;
import com.azure.resourcemanager.compute.models.Sku;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for CapacityReservations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/capacityReservationExamples/CapacityReservation_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a capacity reservation .
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateACapacityReservation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservations()
            .createOrUpdate(
                "myResourceGroup",
                "myCapacityReservationGroup",
                "myCapacityReservation",
                new CapacityReservationInner()
                    .withLocation("westus")
                    .withTags(mapOf("department", "HR"))
                    .withSku(new Sku().withName("Standard_DS1_v2").withCapacity(4L))
                    .withZones(Arrays.asList("1")),
                Context.NONE);
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
