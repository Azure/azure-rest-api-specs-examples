Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.Direction;
import com.azure.resourcemanager.network.models.IpFlowProtocol;
import com.azure.resourcemanager.network.models.VerificationIpFlowParameters;

/** Samples for NetworkWatchers VerifyIpFlow. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherIpFlowVerify.json
     */
    /**
     * Sample code: Ip flow verify.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ipFlowVerify(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .verifyIpFlow(
                "rg1",
                "nw1",
                new VerificationIpFlowParameters()
                    .withTargetResourceId(
                        "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1")
                    .withDirection(Direction.OUTBOUND)
                    .withProtocol(IpFlowProtocol.TCP)
                    .withLocalPort("80")
                    .withRemotePort("80")
                    .withLocalIpAddress("10.2.0.4")
                    .withRemoteIpAddress("121.10.1.1"),
                Context.NONE);
    }
}
```
