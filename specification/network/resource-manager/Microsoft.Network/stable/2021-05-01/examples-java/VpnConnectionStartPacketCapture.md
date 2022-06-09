```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.VpnConnectionPacketCaptureStartParameters;
import java.util.Arrays;

/** Samples for VpnConnections StartPacketCapture. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnConnectionStartPacketCapture.json
     */
    /**
     * Sample code: Start packet capture on vpn connection without filter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startPacketCaptureOnVpnConnectionWithoutFilter(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnConnections()
            .startPacketCapture(
                "rg1",
                "gateway1",
                "vpnConnection1",
                new VpnConnectionPacketCaptureStartParameters()
                    .withLinkConnectionNames(Arrays.asList("siteLink1", "siteLink2")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
