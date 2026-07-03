
import com.azure.resourcemanager.network.fluent.models.RoutingRuleCollectionInner;
import com.azure.resourcemanager.network.models.NetworkManagerRoutingGroupItem;
import java.util.Arrays;

/**
 * Samples for RoutingRuleCollections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingRuleCollectionPut.json
     */
    /**
     * Sample code: Create or Update a routing rule collection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createOrUpdateARoutingRuleCollection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutingRuleCollections().createOrUpdateWithResponse("rg1", "testNetworkManager",
            "myTestRoutingConfig", "testRuleCollection",
            new RoutingRuleCollectionInner().withDescription("A sample policy")
                .withAppliesTo(Arrays.asList(new NetworkManagerRoutingGroupItem().withNetworkGroupId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/testGroup"))),
            com.azure.core.util.Context.NONE);
    }
}
