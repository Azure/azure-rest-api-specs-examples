
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;
import com.azure.resourcemanager.network.models.PublicIpAddressDnsSettings;
import com.azure.resourcemanager.network.models.PublicIpAddressDnsSettingsDomainNameLabelScope;

/**
 * Samples for PublicIpAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressCreateDnsWithDomainNameLabelScope.json
     */
    /**
     * Sample code: Create public IP address DNS with Domain Name Label Scope.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createPublicIPAddressDNSWithDomainNameLabelScope(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().createOrUpdate("rg1", "test-ip",
            new PublicIpAddressInner().withLocation("eastus")
                .withDnsSettings(new PublicIpAddressDnsSettings().withDomainNameLabel("dnslbl")
                    .withDomainNameLabelScope(PublicIpAddressDnsSettingsDomainNameLabelScope.TENANT_REUSE)),
            com.azure.core.util.Context.NONE);
    }
}
