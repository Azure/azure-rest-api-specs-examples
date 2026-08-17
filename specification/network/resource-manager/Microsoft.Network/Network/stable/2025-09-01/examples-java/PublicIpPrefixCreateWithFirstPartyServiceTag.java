
import com.azure.resourcemanager.network.fluent.models.PublicIpPrefixInner;
import com.azure.resourcemanager.network.models.IpTag;
import com.azure.resourcemanager.network.models.IpVersion;
import com.azure.resourcemanager.network.models.PublicIpPrefixSku;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuName;
import com.azure.resourcemanager.network.models.PublicIpPrefixSkuTier;
import java.util.Arrays;

/**
 * Samples for PublicIpPrefixes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PublicIpPrefixCreateWithFirstPartyServiceTag.json
     */
    /**
     * Sample code: Create public IP prefix with first party service tag.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createPublicIPPrefixWithFirstPartyServiceTag(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpPrefixes().createOrUpdate("rg1", "test-ipprefix", new PublicIpPrefixInner()
            .withLocation("westus")
            .withSku(new PublicIpPrefixSku()
                .withName(PublicIpPrefixSkuName.STANDARD).withTier(PublicIpPrefixSkuTier.REGIONAL))
            .withPublicIpAddressVersion(IpVersion.IPV4)
            .withIpTags(
                Arrays.asList(new IpTag().withIpTagType("FirstPartyUsage").withTag("SQL").withFirstPartyServiceTagId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firstPartyServiceTags/myServiceTag")))
            .withPrefixLength(30), com.azure.core.util.Context.NONE);
    }
}
