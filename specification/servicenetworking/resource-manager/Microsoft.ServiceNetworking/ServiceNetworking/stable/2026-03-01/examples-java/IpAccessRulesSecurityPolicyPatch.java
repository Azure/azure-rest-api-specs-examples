
import com.azure.resourcemanager.servicenetworking.models.IpAccessRulesPolicy;
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicy;
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicyUpdateProperties;
import java.util.Arrays;

/**
 * Samples for SecurityPoliciesInterface Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/IpAccessRulesSecurityPolicyPatch.json
     */
    /**
     * Sample code: Update IpAccessRules SecurityPolicy.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void updateIpAccessRulesSecurityPolicy(
        com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        SecurityPolicy resource = manager.securityPoliciesInterfaces()
            .getWithResponse("rg1", "tc1", "sp1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new SecurityPolicyUpdateProperties()
            .withIpAccessRulesPolicy(new IpAccessRulesPolicy().withRules(Arrays.asList()))).apply();
    }
}
