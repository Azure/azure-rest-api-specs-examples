Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.PublicIpPrefixInner;
import com.azure.resourcemanager.network.models.PublicIpPrefixSku;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuName;

/** Samples for PublicIpPrefixes CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpPrefixCreateDefaults.json
     */
    /**
     * Sample code: Create public IP prefix defaults.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPublicIPPrefixDefaults(com.azure.resourcemanager.AzureResourceManager azure) {
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
                    .withSku(new PublicIpPrefixSku().withName(PublicIpPrefixSkuName.STANDARD))
                    .withPrefixLength(30),
                Context.NONE);
    }
}
```
