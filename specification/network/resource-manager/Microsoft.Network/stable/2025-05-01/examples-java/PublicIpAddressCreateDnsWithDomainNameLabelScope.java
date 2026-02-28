
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;
import com.azure.resourcemanager.network.models.PublicIpAddressDnsSettings;
import com.azure.resourcemanager.network.models.PublicIpAddressDnsSettingsDomainNameLabelScope;

/**
 * Samples for PublicIpAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * PublicIpAddressCreateDnsWithDomainNameLabelScope.json
     */
    /**
     * Sample code: Create public IP address DNS with Domain Name Label Scope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createPublicIPAddressDNSWithDomainNameLabelScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().createOrUpdate("rg1", "test-ip",
            new PublicIpAddressInner().withLocation("eastus")
                .withDnsSettings(new PublicIpAddressDnsSettings().withDomainNameLabel("dnslbl")
                    .withDomainNameLabelScope(PublicIpAddressDnsSettingsDomainNameLabelScope.TENANT_REUSE)),
            com.azure.core.util.Context.NONE);
    }
}
