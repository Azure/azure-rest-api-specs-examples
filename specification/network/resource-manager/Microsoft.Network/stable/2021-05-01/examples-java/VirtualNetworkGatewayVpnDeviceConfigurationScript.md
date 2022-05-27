Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.VpnDeviceScriptParameters;

/** Samples for VirtualNetworkGateways VpnDeviceConfigurationScript. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayVpnDeviceConfigurationScript.json
     */
    /**
     * Sample code: GetVPNDeviceConfigurationScript.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVPNDeviceConfigurationScript(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .vpnDeviceConfigurationScriptWithResponse(
                "rg1",
                "vpngw",
                new VpnDeviceScriptParameters()
                    .withVendor("Cisco")
                    .withDeviceFamily("ISR")
                    .withFirmwareVersion("IOS 15.1 (Preview)"),
                Context.NONE);
    }
}
```
