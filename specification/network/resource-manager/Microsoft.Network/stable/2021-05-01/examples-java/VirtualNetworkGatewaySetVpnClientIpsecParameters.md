```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VpnClientIPsecParametersInner;
import com.azure.resourcemanager.network.models.DhGroup;
import com.azure.resourcemanager.network.models.IkeEncryption;
import com.azure.resourcemanager.network.models.IkeIntegrity;
import com.azure.resourcemanager.network.models.IpsecEncryption;
import com.azure.resourcemanager.network.models.IpsecIntegrity;
import com.azure.resourcemanager.network.models.PfsGroup;

/** Samples for VirtualNetworkGateways SetVpnclientIpsecParameters. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewaySetVpnClientIpsecParameters.json
     */
    /**
     * Sample code: Set VirtualNetworkGateway VpnClientIpsecParameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void setVirtualNetworkGatewayVpnClientIpsecParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .setVpnclientIpsecParameters(
                "rg1",
                "vpngw",
                new VpnClientIPsecParametersInner()
                    .withSaLifeTimeSeconds(86473)
                    .withSaDataSizeKilobytes(429497)
                    .withIpsecEncryption(IpsecEncryption.AES256)
                    .withIpsecIntegrity(IpsecIntegrity.SHA256)
                    .withIkeEncryption(IkeEncryption.AES256)
                    .withIkeIntegrity(IkeIntegrity.SHA384)
                    .withDhGroup(DhGroup.DHGROUP2)
                    .withPfsGroup(PfsGroup.PFS2),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
