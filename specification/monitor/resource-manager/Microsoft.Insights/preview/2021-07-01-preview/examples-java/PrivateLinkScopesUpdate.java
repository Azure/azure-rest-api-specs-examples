
import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.fluent.models.AzureMonitorPrivateLinkScopeInner;
import com.azure.resourcemanager.monitor.models.AccessMode;
import com.azure.resourcemanager.monitor.models.AccessModeSettings;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for PrivateLinkScopes CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateLinkScopesUpdate.json
     */
    /**
     * Sample code: PrivateLinkScopeUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateLinkScopeUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateLinkScopes().createOrUpdateWithResponse(
            "my-resource-group", "my-privatelinkscope",
            new AzureMonitorPrivateLinkScopeInner().withLocation("Global").withTags(mapOf("Tag1", "Value1"))
                .withAccessModeSettings(new AccessModeSettings().withQueryAccessMode(AccessMode.OPEN)
                    .withIngestionAccessMode(AccessMode.OPEN).withExclusions(Arrays.asList())),
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
