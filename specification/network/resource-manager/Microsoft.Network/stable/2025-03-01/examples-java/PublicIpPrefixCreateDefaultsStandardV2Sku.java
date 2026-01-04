
import com.azure.resourcemanager.network.fluent.models.PublicIpPrefixInner;
import com.azure.resourcemanager.network.models.PublicIpPrefixSku;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuName;

/**
 * Samples for PublicIpPrefixes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * PublicIpPrefixCreateDefaultsStandardV2Sku.json
     */
    /**
     * Sample code: Create public IP prefix defaults with StandardV2 Sku.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createPublicIPPrefixDefaultsWithStandardV2Sku(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpPrefixes().createOrUpdate("rg1", "test-ipprefix",
            new PublicIpPrefixInner().withLocation("westus")
                .withSku(new PublicIpPrefixSku().withName(PublicIpPrefixSkuName.STANDARD_V2)).withPrefixLength(30),
            com.azure.core.util.Context.NONE);
    }
}
