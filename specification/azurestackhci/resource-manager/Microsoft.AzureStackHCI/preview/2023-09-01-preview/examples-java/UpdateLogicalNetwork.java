import com.azure.resourcemanager.azurestackhci.models.LogicalNetworks;
import java.util.HashMap;
import java.util.Map;

/** Samples for LogicalNetworksOperation Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateLogicalNetwork.json
     */
    /**
     * Sample code: UpdateLogicalNetwork.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void updateLogicalNetwork(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        LogicalNetworks resource =
            manager
                .logicalNetworksOperations()
                .getByResourceGroupWithResponse("test-rg", "test-lnet", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("additionalProperties", "sample")).apply();
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
