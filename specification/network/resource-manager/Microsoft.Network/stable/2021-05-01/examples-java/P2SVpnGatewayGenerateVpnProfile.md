Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.AuthenticationMethod;
import com.azure.resourcemanager.network.models.P2SVpnProfileParameters;

/** Samples for P2SVpnGateways GenerateVpnProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/P2SVpnGatewayGenerateVpnProfile.json
     */
    /**
     * Sample code: GenerateP2SVpnGatewayVPNProfile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateP2SVpnGatewayVPNProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getP2SVpnGateways()
            .generateVpnProfile(
                "rg1",
                "p2sVpnGateway1",
                new P2SVpnProfileParameters().withAuthenticationMethod(AuthenticationMethod.EAPTLS),
                Context.NONE);
    }
}
```
