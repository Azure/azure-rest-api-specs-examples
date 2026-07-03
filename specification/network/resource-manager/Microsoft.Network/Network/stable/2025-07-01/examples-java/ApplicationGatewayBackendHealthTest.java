
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.models.ApplicationGatewayOnDemandProbe;
import com.azure.resourcemanager.network.models.ApplicationGatewayProtocol;

/**
 * Samples for ApplicationGateways BackendHealthOnDemand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayBackendHealthTest.json
     */
    /**
     * Sample code: Test Backend Health.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void testBackendHealth(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().backendHealthOnDemand("rg1", "appgw",
            new ApplicationGatewayOnDemandProbe().withProtocol(ApplicationGatewayProtocol.HTTP).withPath("/")
                .withTimeout(30).withPickHostnameFromBackendHttpSettings(true)
                .withBackendAddressPool(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendaddressPools/MFAnalyticsPool"))
                .withBackendHttpSettings(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/MFPoolSettings")),
            null, com.azure.core.util.Context.NONE);
    }
}
