
import com.azure.resourcemanager.deviceupdate.fluent.models.PrivateEndpointConnectionProxyInner;
import com.azure.resourcemanager.deviceupdate.models.PrivateLinkServiceConnection;
import com.azure.resourcemanager.deviceupdate.models.PrivateLinkServiceProxy;
import com.azure.resourcemanager.deviceupdate.models.RemotePrivateEndpoint;
import java.util.Arrays;

/**
 * Samples for PrivateEndpointConnectionProxies Validate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/
     * PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Validate.json
     */
    /**
     * Sample code: PrivateEndpointConnectionProxyValidate.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void
        privateEndpointConnectionProxyValidate(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnectionProxies().validateWithResponse("test-rg", "contoso", "peexample01",
            new PrivateEndpointConnectionProxyInner().withRemotePrivateEndpoint(new RemotePrivateEndpoint().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}")
                .withLocation("westus2").withImmutableSubscriptionId("00000000-0000-0000-0000-000000000000")
                .withImmutableResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}")
                .withManualPrivateLinkServiceConnections(Arrays.asList(new PrivateLinkServiceConnection()
                    .withName("{privateEndpointConnectionProxyId}").withGroupIds(Arrays.asList("DeviceUpdate"))
                    .withRequestMessage("Please approve my connection, thanks.")))
                .withPrivateLinkServiceProxies(Arrays.asList(new PrivateLinkServiceProxy().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}")
                    .withGroupConnectivityInformation(Arrays.asList())))),
            com.azure.core.util.Context.NONE);
    }
}
