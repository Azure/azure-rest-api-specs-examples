import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.IpGroupInner;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for IpGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/IpGroupsCreate.json
     */
    /**
     * Sample code: CreateOrUpdate_IpGroups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateIpGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getIpGroups()
            .createOrUpdate(
                "myResourceGroup",
                "ipGroups1",
                new IpGroupInner()
                    .withLocation("West US")
                    .withTags(mapOf("key1", "value1"))
                    .withIpAddresses(Arrays.asList("13.64.39.16/32", "40.74.146.80/31", "40.74.147.32/28")),
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
