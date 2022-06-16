import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.CapacityReservationGroupInner;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for CapacityReservationGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/capacityReservationExamples/CapacityReservationGroup_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a capacity reservation group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateACapacityReservationGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservationGroups()
            .createOrUpdateWithResponse(
                "myResourceGroup",
                "myCapacityReservationGroup",
                new CapacityReservationGroupInner()
                    .withLocation("westus")
                    .withTags(mapOf("department", "finance"))
                    .withZones(Arrays.asList("1", "2")),
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
