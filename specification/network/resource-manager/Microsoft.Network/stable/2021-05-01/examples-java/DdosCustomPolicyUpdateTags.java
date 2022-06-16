import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.TagsObject;
import java.util.HashMap;
import java.util.Map;

/** Samples for DdosCustomPolicies UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DdosCustomPolicyUpdateTags.json
     */
    /**
     * Sample code: DDoS Custom policy Update tags.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dDoSCustomPolicyUpdateTags(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getDdosCustomPolicies()
            .updateTagsWithResponse(
                "rg1",
                "test-ddos-custom-policy",
                new TagsObject().withTags(mapOf("tag1", "value1", "tag2", "value2")),
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
