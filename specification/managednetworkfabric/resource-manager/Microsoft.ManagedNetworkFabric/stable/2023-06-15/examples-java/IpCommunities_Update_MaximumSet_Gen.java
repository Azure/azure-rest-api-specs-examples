import com.azure.resourcemanager.managednetworkfabric.models.CommunityActionTypes;
import com.azure.resourcemanager.managednetworkfabric.models.IpCommunity;
import com.azure.resourcemanager.managednetworkfabric.models.IpCommunityRule;
import com.azure.resourcemanager.managednetworkfabric.models.WellKnownCommunities;
import java.util.Arrays;

/** Samples for IpCommunities Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpCommunities_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpCommunities_Update_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipCommunitiesUpdateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        IpCommunity resource =
            manager
                .ipCommunities()
                .getByResourceGroupWithResponse("example-rg", "example-ipcommunity", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withIpCommunityRules(
                Arrays
                    .asList(
                        new IpCommunityRule()
                            .withAction(CommunityActionTypes.PERMIT)
                            .withSequenceNumber(4155123341L)
                            .withWellKnownCommunities(Arrays.asList(WellKnownCommunities.INTERNET))
                            .withCommunityMembers(Arrays.asList("1:1"))))
            .apply();
    }
}
