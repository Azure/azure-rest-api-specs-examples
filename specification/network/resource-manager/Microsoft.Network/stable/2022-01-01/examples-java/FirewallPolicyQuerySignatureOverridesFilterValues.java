import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.SignatureOverridesFilterValuesQuery;

/** Samples for FirewallPolicyIdpsSignaturesFilterValues List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/FirewallPolicyQuerySignatureOverridesFilterValues.json
     */
    /**
     * Sample code: query signature overrides.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void querySignatureOverrides(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyIdpsSignaturesFilterValues()
            .listWithResponse(
                "rg1",
                "firewallPolicy",
                new SignatureOverridesFilterValuesQuery().withFilterName("severity"),
                Context.NONE);
    }
}
