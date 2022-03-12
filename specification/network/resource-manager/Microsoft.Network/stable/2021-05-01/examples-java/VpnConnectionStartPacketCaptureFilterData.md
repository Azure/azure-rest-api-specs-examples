Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.VpnConnectionPacketCaptureStartParameters;
import java.util.Arrays;

/** Samples for VpnConnections StartPacketCapture. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnConnectionStartPacketCaptureFilterData.json
     */
    /**
     * Sample code: Start packet capture on vpn connection with filter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startPacketCaptureOnVpnConnectionWithFilter(
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
                    .withFilterData(
                        "{'TracingFlags': 11,'MaxPacketBufferSize': 120,'MaxFileSize': 200,'Filters':"
                            + " [{'SourceSubnets': ['20.1.1.0/24'],'DestinationSubnets': ['10.1.1.0/24'],'SourcePort':"
                            + " [500],'DestinationPort': [4500],'Protocol': 6,'TcpFlags':"
                            + " 16,'CaptureSingleDirectionTrafficOnly': true}]}")
                    .withLinkConnectionNames(Arrays.asList("siteLink1", "siteLink2")),
                Context.NONE);
    }
}
```
