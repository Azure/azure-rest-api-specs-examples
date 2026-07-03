
import com.azure.resourcemanager.network.fluent.models.NspAccessRuleInner;
import com.azure.resourcemanager.network.models.AccessRuleDirection;
import java.util.Arrays;

/**
 * Samples for NetworkSecurityPerimeterAccessRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspAccessRulePut.json
     */
    /**
     * Sample code: NspAccessRulePut.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspAccessRulePut(com.azure.resourcemanager.network.NetworkManager manager) {
        manager
            .serviceClient().getNetworkSecurityPerimeterAccessRules().createOrUpdateWithResponse("rg1", "nsp1",
                "profile1", "accessRule1", new NspAccessRuleInner().withDirection(AccessRuleDirection.INBOUND)
                    .withAddressPrefixes(Arrays.asList("10.11.0.0/16", "10.10.1.0/24")),
                com.azure.core.util.Context.NONE);
    }
}
