import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.CdnWebApplicationFirewallPolicyPatchParameters;
import java.util.HashMap;
import java.util.Map;

/** Samples for Policies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPatchPolicy.json
     */
    /**
     * Sample code: Creates specific policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsSpecificPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getPolicies()
            .update(
                "rg1",
                "MicrosoftCdnWafPolicy",
                new CdnWebApplicationFirewallPolicyPatchParameters().withTags(mapOf("foo", "bar")),
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
