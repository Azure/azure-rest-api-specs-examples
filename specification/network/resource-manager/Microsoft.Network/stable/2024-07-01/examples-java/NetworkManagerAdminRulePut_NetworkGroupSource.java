
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
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * NetworkManagerAdminRulePut_NetworkGroupSource.json
     */
    /**
     * Sample code: Create a admin rule with network group as source or destination.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAAdminRuleWithNetworkGroupAsSourceOrDestination(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAdminRules().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleAdminRule",
            new AdminRule().withDescription("This is Sample Admin Rule")
                .withProtocol(SecurityConfigurationRuleProtocol.TCP)
                .withSources(Arrays.asList(new AddressPrefixItem().withAddressPrefix("Internet")
                    .withAddressPrefixType(AddressPrefixType.SERVICE_TAG)))
                .withDestinations(Arrays.asList(new AddressPrefixItem().withAddressPrefix(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/ng1")
                    .withAddressPrefixType(AddressPrefixType.NETWORK_GROUP)))
                .withSourcePortRanges(Arrays.asList("0-65535")).withDestinationPortRanges(Arrays.asList("22"))
                .withAccess(SecurityConfigurationRuleAccess.DENY).withPriority(1)
                .withDirection(SecurityConfigurationRuleDirection.INBOUND),
            com.azure.core.util.Context.NONE);
    }
}
