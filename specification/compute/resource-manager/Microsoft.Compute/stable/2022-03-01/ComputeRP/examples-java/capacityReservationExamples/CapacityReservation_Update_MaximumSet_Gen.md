Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.CapacityReservationUpdate;
import com.azure.resourcemanager.compute.models.Sku;
import java.util.HashMap;
import java.util.Map;

/** Samples for CapacityReservations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/capacityReservationExamples/CapacityReservation_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservations_Update_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void capacityReservationsUpdateMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservations()
            .update(
                "rgcompute",
                "aaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaa",
                new CapacityReservationUpdate()
                    .withTags(mapOf("key4974", "aaaaaaaaaaaaaaaa"))
                    .withSku(new Sku().withName("DSv3-Type1").withTier("aaa").withCapacity(7L)),
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
```
