
import com.azure.resourcemanager.network.models.AddressPrefixItem;
import com.azure.resourcemanager.network.models.AddressPrefixType;
import com.azure.resourcemanager.network.models.AdminRule;
import com.azure.resourcemanager.network.models.SecurityConfigurationRuleAccess;
import com.azure.resourcemanager.network.models.SecurityConfigurationRuleDirection;
import com.azure.resourcemanager.network.models.SecurityConfigurationRuleProtocol;
import java.util.Arrays;

/**
 * Samples for AdminRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerAdminRulePut.
     * json
     */
    /**
     * Sample code: Create an admin rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAnAdminRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAdminRules().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleAdminRule",
            new AdminRule().withDescription("This is Sample Admin Rule")
                .withProtocol(SecurityConfigurationRuleProtocol.TCP)
                .withSources(Arrays.asList(new AddressPrefixItem().withAddressPrefix("Internet")
                    .withAddressPrefixType(AddressPrefixType.SERVICE_TAG)))
                .withDestinations(Arrays.asList(
                    new AddressPrefixItem().withAddressPrefix("*").withAddressPrefixType(AddressPrefixType.IPPREFIX)))
                .withSourcePortRanges(Arrays.asList("0-65535")).withDestinationPortRanges(Arrays.asList("22"))
                .withAccess(SecurityConfigurationRuleAccess.DENY).withPriority(1)
                .withDirection(SecurityConfigurationRuleDirection.INBOUND),
            com.azure.core.util.Context.NONE);
    }
}
