Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceupdate_1.0.0-beta.1/sdk/deviceupdate/azure-resourcemanager-deviceupdate/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.deviceupdate.models.PrivateEndpointUpdate;

/** Samples for PrivateEndpointConnectionProxies UpdatePrivateEndpointProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_PrivateEndpointUpdate.json
     */
    /**
     * Sample code: PrivateEndpointConnectionProxyPrivateEndpointUpdate.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateEndpointConnectionProxyPrivateEndpointUpdate(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager
            .privateEndpointConnectionProxies()
            .updatePrivateEndpointPropertiesWithResponse(
                "test-rg",
                "contoso",
                "peexample01",
                new PrivateEndpointUpdate()
                    .withId(
                        "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}")
                    .withLocation("westus2")
                    .withImmutableSubscriptionId("00000000-0000-0000-0000-000000000000")
                    .withImmutableResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}")
                    .withVnetTrafficTag("12345678"),
                Context.NONE);
    }
}
```
