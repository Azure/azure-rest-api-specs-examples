import com.azure.resourcemanager.managednetworkfabric.models.ExpressRouteConnectionInformation;
import com.azure.resourcemanager.managednetworkfabric.models.IsWorkloadManagementNetworkEnabled;
import com.azure.resourcemanager.managednetworkfabric.models.ManagedResourceGroupConfiguration;
import com.azure.resourcemanager.managednetworkfabric.models.NfcSku;
import java.util.Arrays;

/** Samples for NetworkFabricControllers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabricControllers_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabricControllers_Create_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricControllersCreateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabricControllers()
            .define("example-networkController")
            .withRegion("eastus")
            .withExistingResourceGroup("example-rg")
            .withManagedResourceGroupConfiguration(
                new ManagedResourceGroupConfiguration().withName("managedResourceGroupName").withLocation("eastus"))
            .withIsWorkloadManagementNetworkEnabled(IsWorkloadManagementNetworkEnabled.TRUE)
            .withIpv4AddressSpace("172.253.0.0/19")
            .withIpv6AddressSpace("::/60")
            .withNfcSku(NfcSku.STANDARD)
            .withInfrastructureExpressRouteConnections(
                Arrays
                    .asList(
                        new ExpressRouteConnectionInformation()
                            .withExpressRouteCircuitId(
                                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName")
                            .withExpressRouteAuthorizationKey("fakeTokenPlaceholder")))
            .withWorkloadExpressRouteConnections(
                Arrays
                    .asList(
                        new ExpressRouteConnectionInformation()
                            .withExpressRouteCircuitId(
                                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName")
                            .withExpressRouteAuthorizationKey("fakeTokenPlaceholder")))
            .withAnnotation("annotation")
            .create();
    }
}
