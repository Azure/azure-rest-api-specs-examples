
import com.azure.resourcemanager.network.fluent.models.SecurityRuleInner;
import com.azure.resourcemanager.network.models.SecurityRuleAccess;
import com.azure.resourcemanager.network.models.SecurityRuleDirection;
import com.azure.resourcemanager.network.models.SecurityRuleProtocol;

/**
 * Samples for SecurityRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupRuleCreate.json
     */
    /**
     * Sample code: Create security rule.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createSecurityRule(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityRules().createOrUpdate("rg1", "testnsg", "rule1",
            new SecurityRuleInner().withProtocol(SecurityRuleProtocol.ASTERISK).withSourcePortRange("*")
                .withDestinationPortRange("8080").withSourceAddressPrefix("10.0.0.0/8")
                .withDestinationAddressPrefix("11.0.0.0/8").withAccess(SecurityRuleAccess.DENY).withPriority(100)
                .withDirection(SecurityRuleDirection.OUTBOUND),
            com.azure.core.util.Context.NONE);
    }
}
