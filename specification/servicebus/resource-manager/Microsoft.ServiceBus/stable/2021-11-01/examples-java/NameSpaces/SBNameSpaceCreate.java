
import com.azure.resourcemanager.servicebus.fluent.models.SBNamespaceInner;
import com.azure.resourcemanager.servicebus.models.SBSku;
import com.azure.resourcemanager.servicebus.models.SkuName;
import com.azure.resourcemanager.servicebus.models.SkuTier;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Namespaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * SBNameSpaceCreate.json
     */
    /**
     * Sample code: NameSpaceCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces().createOrUpdate("ArunMonocle",
            "sdk-Namespace2924",
            new SBNamespaceInner().withLocation("South Central US").withTags(mapOf("tag1", "value1", "tag2", "value2"))
                .withSku(new SBSku().withName(SkuName.STANDARD).withTier(SkuTier.STANDARD)),
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
