
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;
import com.azure.resourcemanager.network.models.IpAllocationMethod;
import com.azure.resourcemanager.network.models.IpVersion;
import com.azure.resourcemanager.network.models.PublicIpAddressSku;
import com.azure.resourcemanager.network.models.PublicIpAddressSkuName;
import com.azure.resourcemanager.network.models.PublicIpAddressSkuTier;

/**
 * Samples for PublicIpAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressCreateCustomizedValues.json
     */
    /**
     * Sample code: Create public IP address allocation method.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createPublicIPAddressAllocationMethod(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().createOrUpdate("rg1", "test-ip",
            new PublicIpAddressInner().withLocation("eastus")
                .withSku(new PublicIpAddressSku().withName(PublicIpAddressSkuName.STANDARD)
                    .withTier(PublicIpAddressSkuTier.GLOBAL))
                .withPublicIpAllocationMethod(IpAllocationMethod.STATIC).withPublicIpAddressVersion(IpVersion.IPV4)
                .withIdleTimeoutInMinutes(10),
            com.azure.core.util.Context.NONE);
    }
}
