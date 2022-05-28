Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.ApplicationGatewayOnDemandProbe;
import com.azure.resourcemanager.network.models.ApplicationGatewayProtocol;

/** Samples for ApplicationGateways BackendHealthOnDemand. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayBackendHealthTest.json
     */
    /**
     * Sample code: Test Backend Health.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void testBackendHealth(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationGateways()
            .backendHealthOnDemand(
                "rg1",
                "appgw",
                new ApplicationGatewayOnDemandProbe()
                    .withProtocol(ApplicationGatewayProtocol.HTTP)
                    .withPath("/")
                    .withTimeout(30)
                    .withPickHostnameFromBackendHttpSettings(true)
                    .withBackendAddressPool(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendaddressPools/MFAnalyticsPool"))
                    .withBackendHttpSettings(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/MFPoolSettings")),
                null,
                Context.NONE);
    }
}
```
