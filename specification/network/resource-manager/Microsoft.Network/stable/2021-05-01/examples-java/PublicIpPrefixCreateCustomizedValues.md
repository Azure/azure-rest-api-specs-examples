Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.PublicIpPrefixInner;
import com.azure.resourcemanager.network.models.IpVersion;
import com.azure.resourcemanager.network.models.PublicIpPrefixSku;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuName;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuTier;

/** Samples for PublicIpPrefixes CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpPrefixCreateCustomizedValues.json
     */
    /**
     * Sample code: Create public IP prefix allocation method.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPublicIPPrefixAllocationMethod(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPublicIpPrefixes()
            .createOrUpdate(
                "rg1",
                "test-ipprefix",
                new PublicIpPrefixInner()
                    .withLocation("westus")
                    .withSku(
                        new PublicIpPrefixSku()
                            .withName(PublicIpPrefixSkuName.STANDARD)
                            .withTier(PublicIpPrefixSkuTier.REGIONAL))
                    .withPublicIpAddressVersion(IpVersion.IPV4)
                    .withPrefixLength(30),
                Context.NONE);
    }
}
```
