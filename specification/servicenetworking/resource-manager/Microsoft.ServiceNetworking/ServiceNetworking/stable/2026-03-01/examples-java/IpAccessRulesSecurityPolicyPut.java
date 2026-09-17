
import com.azure.resourcemanager.servicenetworking.models.IpAccessRulesPolicy;
import com.azure.resourcemanager.servicenetworking.models.SecurityPolicyProperties;
import java.util.Arrays;

/**
 * Samples for SecurityPoliciesInterface CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/IpAccessRulesSecurityPolicyPut.json
     */
    /**
     * Sample code: Put IpAccessRules SecurityPolicy.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        putIpAccessRulesSecurityPolicy(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.securityPoliciesInterfaces().define("sp1").withRegion("NorthCentralUS")
            .withExistingTrafficController("rg1", "tc1").withProperties(new SecurityPolicyProperties()
                .withIpAccessRulesPolicy(new IpAccessRulesPolicy().withRules(Arrays.asList())))
            .create();
    }
}
