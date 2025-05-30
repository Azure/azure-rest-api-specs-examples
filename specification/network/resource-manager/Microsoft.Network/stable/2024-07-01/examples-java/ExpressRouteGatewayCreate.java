
import com.azure.resourcemanager.network.fluent.models.ExpressRouteGatewayInner;
import com.azure.resourcemanager.network.models.ExpressRouteGatewayPropertiesAutoScaleConfiguration;
import com.azure.resourcemanager.network.models.ExpressRouteGatewayPropertiesAutoScaleConfigurationBounds;
import com.azure.resourcemanager.network.models.VirtualHubId;

/**
 * Samples for ExpressRouteGateways CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ExpressRouteGatewayCreate.
     * json
     */
    /**
     * Sample code: ExpressRouteGatewayCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteGatewayCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteGateways()
            .createOrUpdate("resourceGroupName", "gateway-2", new ExpressRouteGatewayInner().withLocation("westus")
                .withAutoScaleConfiguration(new ExpressRouteGatewayPropertiesAutoScaleConfiguration()
                    .withBounds(new ExpressRouteGatewayPropertiesAutoScaleConfigurationBounds().withMin(3)))
                .withVirtualHub(new VirtualHubId().withId(
                    "/subscriptions/subid/resourceGroups/resourceGroupId/providers/Microsoft.Network/virtualHubs/virtualHubName"))
                .withAllowNonVirtualWanTraffic(false), com.azure.core.util.Context.NONE);
    }
}
