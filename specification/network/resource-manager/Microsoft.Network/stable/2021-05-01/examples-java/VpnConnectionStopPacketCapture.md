Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.VpnConnectionPacketCaptureStopParameters;
import java.util.Arrays;

/** Samples for VpnConnections StopPacketCapture. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnConnectionStopPacketCapture.json
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
            .stopPacketCapture(
                "rg1",
                "gateway1",
                "vpnConnection1",
                new VpnConnectionPacketCaptureStopParameters()
                    .withSasUrl(
                        "https://teststorage.blob.core.windows.net/?sv=2018-03-28&ss=bfqt&srt=sco&sp=rwdlacup&se=2019-09-13T07:44:05Z&st=2019-09-06T23:44:05Z&spr=https&sig=V1h9D1riltvZMI69d6ihENnFo%2FrCvTqGgjO2lf%2FVBhE%3D")
                    .withLinkConnectionNames(Arrays.asList("vpnSiteLink1", "vpnSiteLink2")),
                Context.NONE);
    }
}
```
