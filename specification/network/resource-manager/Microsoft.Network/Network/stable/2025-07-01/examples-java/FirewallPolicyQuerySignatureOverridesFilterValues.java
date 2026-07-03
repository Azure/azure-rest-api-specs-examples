
import com.azure.resourcemanager.network.models.SignatureOverridesFilterValuesQuery;

/**
 * Samples for FirewallPolicyIdpsSignaturesFilterValues List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyQuerySignatureOverridesFilterValues.json
     */
    /**
     * Sample code: query signature overrides.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void querySignatureOverrides(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyIdpsSignaturesFilterValues().listWithResponse("rg1", "firewallPolicy",
            new SignatureOverridesFilterValuesQuery().withFilterName("severity"), com.azure.core.util.Context.NONE);
    }
}
