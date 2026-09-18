
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicy;
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicyUpdateProperties;
import com.azure.resourcemanager.servicenetworking.models.WafPolicy;

/**
 * Samples for SecurityPoliciesInterface Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/WafSecurityPolicyPatch.json
     */
    /**
     * Sample code: Update WAF SecurityPolicy.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        updateWAFSecurityPolicy(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        SecurityPolicy resource = manager.securityPoliciesInterfaces()
            .getWithResponse("rg1", "tc1", "sp1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new SecurityPolicyUpdateProperties().withWafPolicy(new WafPolicy().withId(
            "/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/applicationGatewayWebApplicationFirewallPolicies/wp-0")))
            .apply();
    }
}
