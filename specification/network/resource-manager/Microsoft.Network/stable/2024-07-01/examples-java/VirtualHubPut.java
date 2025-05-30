
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.VirtualHubInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualHubs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualHubPut.json
     */
    /**
     * Sample code: VirtualHubPut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubs().createOrUpdate("rg1", "virtualHub2",
            new VirtualHubInner().withLocation("West US").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withVirtualWan(new SubResource().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"))
                .withAddressPrefix("10.168.0.0/24").withSku("Basic"),
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
