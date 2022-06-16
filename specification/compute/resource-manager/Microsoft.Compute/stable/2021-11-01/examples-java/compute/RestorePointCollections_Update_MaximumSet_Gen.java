import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.RestorePointCollectionSourceProperties;
import com.azure.resourcemanager.compute.models.RestorePointCollectionUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for RestorePointCollections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/RestorePointCollections_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: RestorePointCollections_Update_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restorePointCollectionsUpdateMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePointCollections()
            .updateWithResponse(
                "rgcompute",
                "aaaaaaaaaaaaaaaaaaaa",
                new RestorePointCollectionUpdate()
                    .withTags(mapOf("key8536", "aaaaaaaaaaaaaaaaaaa"))
                    .withSource(
                        new RestorePointCollectionSourceProperties()
                            .withId(
                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM")),
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
