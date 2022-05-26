Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceupdate_1.0.0-beta.1/sdk/deviceupdate/azure-resourcemanager-deviceupdate/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.deviceupdate.models.PrivateLinkServiceConnection;
import com.azure.resourcemanager.deviceupdate.models.PrivateLinkServiceProxy;
import com.azure.resourcemanager.deviceupdate.models.RemotePrivateEndpoint;
import java.util.Arrays;

/** Samples for PrivateEndpointConnectionProxies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_CreateOrUpdate.json
     */
    /**
     * Sample code: PrivateEndpointConnectionProxyCreateOrUpdate.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateEndpointConnectionProxyCreateOrUpdate(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager
            .privateEndpointConnectionProxies()
            .define("peexample01")
            .withExistingAccount("test-rg", "contoso")
            .withRemotePrivateEndpoint(
                new RemotePrivateEndpoint()
                    .withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}")
                    .withLocation("westus2")
                    .withImmutableSubscriptionId("00000000-0000-0000-0000-000000000000")
                    .withImmutableResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}")
                    .withManualPrivateLinkServiceConnections(
                        Arrays
                            .asList(
                                new PrivateLinkServiceConnection()
                                    .withName("{privateEndpointConnectionProxyId}")
                                    .withGroupIds(Arrays.asList("DeviceUpdate"))
                                    .withRequestMessage("Please approve my connection, thanks.")))
                    .withPrivateLinkServiceProxies(
                        Arrays
                            .asList(
                                new PrivateLinkServiceProxy()
                                    .withId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}")
                                    .withGroupConnectivityInformation(Arrays.asList()))))
            .create();
    }
}
```
