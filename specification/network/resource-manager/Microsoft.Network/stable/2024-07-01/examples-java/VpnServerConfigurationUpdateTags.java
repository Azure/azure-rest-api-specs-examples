
import com.azure.resourcemanager.network.models.TagsObject;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VpnServerConfigurations UpdateTags.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * VpnServerConfigurationUpdateTags.json
     */
    /**
     * Sample code: VpnServerConfigurationUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnServerConfigurationUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnServerConfigurations().updateTagsWithResponse("rg1",
            "vpnServerConfiguration1",
            new TagsObject().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder")),
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
