import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.AdminRuleCollectionInner;
import com.azure.resourcemanager.network.models.NetworkManagerSecurityGroupItem;
import java.util.Arrays;

/** Samples for AdminRuleCollections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerAdminRuleCollectionPut.json
     */
    /**
     * Sample code: Create or Update an admin rule collection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAnAdminRuleCollection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getAdminRuleCollections()
            .createOrUpdateWithResponse(
                "rg1",
                "testNetworkManager",
                "myTestSecurityConfig",
                "testRuleCollection",
                new AdminRuleCollectionInner()
                    .withDescription("A sample policy")
                    .withAppliesToGroups(
                        Arrays
                            .asList(
                                new NetworkManagerSecurityGroupItem()
                                    .withNetworkGroupId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/testGroup"))),
                Context.NONE);
    }
}
