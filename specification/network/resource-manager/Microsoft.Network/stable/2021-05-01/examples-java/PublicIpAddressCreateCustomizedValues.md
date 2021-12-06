Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;
import com.azure.resourcemanager.network.models.IpAllocationMethod;
import com.azure.resourcemanager.network.models.IpVersion;
import com.azure.resourcemanager.network.models.PublicIpAddressSku;
import com.azure.resourcemanager.network.models.PublicIpAddressSkuName;
import com.azure.resourcemanager.network.models.PublicIpAddressSkuTier;

/** Samples for PublicIpAddresses CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpAddressCreateCustomizedValues.json
     */
    /**
     * Sample code: Create public IP address allocation method.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPublicIPAddressAllocationMethod(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPublicIpAddresses()
            .createOrUpdate(
                "rg1",
                "test-ip",
                new PublicIpAddressInner()
                    .withLocation("eastus")
                    .withSku(
                        new PublicIpAddressSku()
                            .withName(PublicIpAddressSkuName.STANDARD)
                            .withTier(PublicIpAddressSkuTier.GLOBAL))
                    .withPublicIpAllocationMethod(IpAllocationMethod.STATIC)
                    .withPublicIpAddressVersion(IpVersion.IPV4)
                    .withIdleTimeoutInMinutes(10),
                Context.NONE);
    }
}
```
