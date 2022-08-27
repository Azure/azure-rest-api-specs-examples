import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ConnectivityConfigurationInner;
import com.azure.resourcemanager.network.models.ConnectivityGroupItem;
import com.azure.resourcemanager.network.models.ConnectivityTopology;
import com.azure.resourcemanager.network.models.DeleteExistingPeering;
import com.azure.resourcemanager.network.models.GroupConnectivity;
import com.azure.resourcemanager.network.models.Hub;
import com.azure.resourcemanager.network.models.IsGlobal;
import com.azure.resourcemanager.network.models.UseHubGateway;
import java.util.Arrays;

/** Samples for ConnectivityConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectivityConfigurationPut.json
     */
    /**
     * Sample code: ConnectivityConfigurationsPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectivityConfigurationsPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getConnectivityConfigurations()
            .createOrUpdateWithResponse(
                "myResourceGroup",
                "testNetworkManager",
                "myTestConnectivityConfig",
                new ConnectivityConfigurationInner()
                    .withDescription("Sample Configuration")
                    .withConnectivityTopology(ConnectivityTopology.HUB_AND_SPOKE)
                    .withHubs(
                        Arrays
                            .asList(
                                new Hub()
                                    .withResourceId(
                                        "subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myTestConnectivityConfig")
                                    .withResourceType("Microsoft.Network/virtualNetworks")))
                    .withIsGlobal(IsGlobal.TRUE)
                    .withAppliesToGroups(
                        Arrays
                            .asList(
                                new ConnectivityGroupItem()
                                    .withNetworkGroupId(
                                        "subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1")
                                    .withUseHubGateway(UseHubGateway.TRUE)
                                    .withIsGlobal(IsGlobal.FALSE)
                                    .withGroupConnectivity(GroupConnectivity.NONE)))
                    .withDeleteExistingPeering(DeleteExistingPeering.TRUE),
                Context.NONE);
    }
}
