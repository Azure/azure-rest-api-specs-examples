
import com.azure.resourcemanager.network.fluent.models.NspAccessRuleInner;
import com.azure.resourcemanager.network.models.AccessRuleDirection;
import java.util.Arrays;

/**
 * Samples for NetworkSecurityPerimeterAccessRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspAccessRulePut.json
     */
    /**
     * Sample code: NspAccessRulePut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspAccessRulePut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAccessRules().createOrUpdateWithResponse(
            "rg1", "nsp1", "profile1", "accessRule1",
            new NspAccessRuleInner().withDirection(AccessRuleDirection.INBOUND)
                .withAddressPrefixes(Arrays.asList("10.11.0.0/16", "10.10.1.0/24")),
            com.azure.core.util.Context.NONE);
    }
}
