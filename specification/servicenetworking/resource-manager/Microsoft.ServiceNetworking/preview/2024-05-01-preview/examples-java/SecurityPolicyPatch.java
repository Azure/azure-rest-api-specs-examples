
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicy;
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicyUpdateProperties;
import com.azure.resourcemanager.servicenetworking.models.WafPolicyUpdate;

/**
 * Samples for SecurityPoliciesInterface Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/preview/2024-05-01-preview/examples/
     * SecurityPolicyPatch.json
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
        resource.update()
            .withProperties(new SecurityPolicyUpdateProperties().withWafPolicy(new WafPolicyUpdate().withId(
                "/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Networking/applicationGatewayWebApplicationFirewallPolicies/wp-0")))
            .apply();
    }
}
