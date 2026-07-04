
import com.azure.resourcemanager.network.fluent.models.NetworkManagerRoutingConfigurationInner;
import com.azure.resourcemanager.network.models.RouteTableUsageMode;

/**
 * Samples for NetworkManagerRoutingConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingConfigurationPut.json
     */
    /**
     * Sample code: Create network manager routing configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createNetworkManagerRoutingConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagerRoutingConfigurations().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "myTestRoutingConfig", new NetworkManagerRoutingConfigurationInner()
                .withDescription("A sample policy").withRouteTableUsageMode(RouteTableUsageMode.MANAGED_ONLY),
            com.azure.core.util.Context.NONE);
    }
}
