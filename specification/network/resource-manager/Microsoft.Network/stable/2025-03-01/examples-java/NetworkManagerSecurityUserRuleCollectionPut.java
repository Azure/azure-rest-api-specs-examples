
import com.azure.resourcemanager.network.fluent.models.SecurityUserRuleCollectionInner;
import com.azure.resourcemanager.network.models.SecurityUserGroupItem;
import java.util.Arrays;

/**
 * Samples for SecurityUserRuleCollections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerSecurityUserRuleCollectionPut.json
     */
    /**
     * Sample code: Create or Update a Security User Rule Collection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASecurityUserRuleCollection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserRuleCollections().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig", "testRuleCollection",
            new SecurityUserRuleCollectionInner().withDescription("A sample policy")
                .withAppliesToGroups(Arrays.asList(new SecurityUserGroupItem().withNetworkGroupId(
                    "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/testGroup"))),
            com.azure.core.util.Context.NONE);
    }
}
