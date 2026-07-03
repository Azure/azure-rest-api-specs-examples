
import com.azure.resourcemanager.network.fluent.models.PublicIpPrefixInner;
import com.azure.resourcemanager.network.models.PublicIpPrefixSku;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuName;

/**
 * Samples for PublicIpPrefixes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpPrefixCreateDefaultsStandardV2Sku.json
     */
    /**
     * Sample code: Create public IP prefix defaults with StandardV2 Sku.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createPublicIPPrefixDefaultsWithStandardV2Sku(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpPrefixes().createOrUpdate("rg1", "test-ipprefix",
            new PublicIpPrefixInner().withLocation("westus")
                .withSku(new PublicIpPrefixSku().withName(PublicIpPrefixSkuName.STANDARD_V2)).withPrefixLength(30),
            com.azure.core.util.Context.NONE);
    }
}
