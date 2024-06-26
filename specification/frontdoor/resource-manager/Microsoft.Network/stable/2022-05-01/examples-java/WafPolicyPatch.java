import com.azure.resourcemanager.frontdoor.models.WebApplicationFirewallPolicy;
import java.util.HashMap;
import java.util.Map;

/** Samples for Policies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2022-05-01/examples/WafPolicyPatch.json
     */
    /**
     * Sample code: Patches specific policy.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void patchesSpecificPolicy(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        WebApplicationFirewallPolicy resource =
            manager
                .policies()
                .getByResourceGroupWithResponse("rg1", "Policy1", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("key1", "value1", "key2", "value2")).apply();
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
