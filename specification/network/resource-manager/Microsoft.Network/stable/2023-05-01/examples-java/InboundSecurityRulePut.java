import com.azure.resourcemanager.network.fluent.models.InboundSecurityRuleInner;
import com.azure.resourcemanager.network.models.InboundSecurityRules;
import com.azure.resourcemanager.network.models.InboundSecurityRulesProtocol;
import java.util.Arrays;

/** Samples for InboundSecurityRuleOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/InboundSecurityRulePut.json
     */
    /**
     * Sample code: Create Network Virtual Appliance Inbound Security Rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkVirtualApplianceInboundSecurityRules(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getInboundSecurityRuleOperations()
            .createOrUpdate(
                "rg1",
                "nva",
                "rule1",
                new InboundSecurityRuleInner()
                    .withRules(
                        Arrays
                            .asList(
                                new InboundSecurityRules()
                                    .withProtocol(InboundSecurityRulesProtocol.TCP)
                                    .withSourceAddressPrefix("50.20.121.5/32")
                                    .withDestinationPortRange(22))),
                com.azure.core.util.Context.NONE);
    }
}
