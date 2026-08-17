
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;
import com.azure.resourcemanager.network.models.IpAllocationMethod;
import com.azure.resourcemanager.network.models.IpTag;
import com.azure.resourcemanager.network.models.IpVersion;
import com.azure.resourcemanager.network.models.PublicIpAddressSku;
import com.azure.resourcemanager.network.models.PublicIpAddressSkuName;
import com.azure.resourcemanager.network.models.PublicIpAddressSkuTier;
import java.util.Arrays;

/**
 * Samples for PublicIpAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PublicIpAddressCreateWithFirstPartyServiceTag.json
     */
    /**
     * Sample code: Create public IP address with first party service tag.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createPublicIPAddressWithFirstPartyServiceTag(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().createOrUpdate("rg1", "test-ip", new PublicIpAddressInner()
            .withLocation("eastus")
            .withSku(new PublicIpAddressSku().withName(PublicIpAddressSkuName.STANDARD)
                .withTier(PublicIpAddressSkuTier.GLOBAL))
            .withPublicIpAllocationMethod(IpAllocationMethod.STATIC).withPublicIpAddressVersion(IpVersion.IPV4)
            .withIpTags(
                Arrays.asList(new IpTag().withIpTagType("FirstPartyUsage").withTag("SQL").withFirstPartyServiceTagId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firstPartyServiceTags/myServiceTag")))
            .withIdleTimeoutInMinutes(10), com.azure.core.util.Context.NONE);
    }
}
