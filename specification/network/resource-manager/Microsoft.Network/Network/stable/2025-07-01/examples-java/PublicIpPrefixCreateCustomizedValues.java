
import com.azure.resourcemanager.network.fluent.models.PublicIpPrefixInner;
import com.azure.resourcemanager.network.models.IpVersion;
import com.azure.resourcemanager.network.models.PublicIpPrefixSku;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuName;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuTier;

/**
 * Samples for PublicIpPrefixes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpPrefixCreateCustomizedValues.json
     */
    /**
     * Sample code: Create public IP prefix allocation method.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createPublicIPPrefixAllocationMethod(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpPrefixes().createOrUpdate("rg1", "test-ipprefix",
            new PublicIpPrefixInner().withLocation("westus")
                .withSku(new PublicIpPrefixSku().withName(PublicIpPrefixSkuName.STANDARD)
                    .withTier(PublicIpPrefixSkuTier.REGIONAL))
                .withPublicIpAddressVersion(IpVersion.IPV4).withPrefixLength(30),
            com.azure.core.util.Context.NONE);
    }
}
