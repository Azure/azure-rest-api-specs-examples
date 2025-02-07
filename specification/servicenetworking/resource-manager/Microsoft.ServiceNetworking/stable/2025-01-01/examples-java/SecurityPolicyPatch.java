
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicy;
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicyUpdateProperties;
import com.azure.resourcemanager.servicenetworking.models.WafPolicy;

/**
 * Samples for SecurityPoliciesInterface Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SecurityPolicyPatch.json
     */
    /**
     * Sample code: Update SecurityPolicy.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        updateSecurityPolicy(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        SecurityPolicy resource = manager.securityPoliciesInterfaces()
            .getWithResponse("rg1", "tc1", "sp1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new SecurityPolicyUpdateProperties().withWafPolicy(new WafPolicy().withId(
            "/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Networking/applicationGatewayWebApplicationFirewallPolicies/wp-0")))
            .apply();
    }
}
