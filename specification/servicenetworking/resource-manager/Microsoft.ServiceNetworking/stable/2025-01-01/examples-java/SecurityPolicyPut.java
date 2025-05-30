
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicyProperties;
import com.azure.resourcemanager.servicenetworking.models.WafPolicy;

/**
 * Samples for SecurityPoliciesInterface CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SecurityPolicyPut.json
     */
    /**
     * Sample code: Put SecurityPolicy.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void putSecurityPolicy(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.securityPoliciesInterfaces().define("sp1").withRegion("NorthCentralUS")
            .withExistingTrafficController("rg1", "tc1")
            .withProperties(new SecurityPolicyProperties().withWafPolicy(new WafPolicy().withId(
                "/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Networking/applicationGatewayWebApplicationFirewallPolicies/wp-0")))
            .create();
    }
}
