
import com.azure.resourcemanager.network.fluent.models.PublicIpPrefixInner;
import com.azure.resourcemanager.network.models.PublicIpPrefixSku;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuName;

/**
 * Samples for PublicIpPrefixes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PublicIpPrefixCreateDefaults.
     * json
     */
    /**
     * Sample code: Create public IP prefix defaults.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPublicIPPrefixDefaults(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpPrefixes().createOrUpdate("rg1", "test-ipprefix",
            new PublicIpPrefixInner().withLocation("westus")
                .withSku(new PublicIpPrefixSku().withName(PublicIpPrefixSkuName.STANDARD)).withPrefixLength(30),
            com.azure.core.util.Context.NONE);
    }
}
