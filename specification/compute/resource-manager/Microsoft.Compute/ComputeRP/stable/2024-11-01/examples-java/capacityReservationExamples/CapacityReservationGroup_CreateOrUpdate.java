
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.fluent.models.CapacityReservationGroupInner;
import com.azure.resourcemanager.compute.models.ResourceSharingProfile;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CapacityReservationGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * capacityReservationExamples/CapacityReservationGroup_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a capacity reservation group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateACapacityReservationGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservationGroups().createOrUpdateWithResponse(
            "myResourceGroup", "myCapacityReservationGroup",
            new CapacityReservationGroupInner().withLocation("westus").withTags(mapOf("department", "finance"))
                .withZones(Arrays.asList("1", "2"))
                .withSharingProfile(new ResourceSharingProfile()
                    .withSubscriptionIds(Arrays.asList(new SubResource().withId("/subscriptions/{subscription-id1}"),
                        new SubResource().withId("/subscriptions/{subscription-id2}")))),
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
