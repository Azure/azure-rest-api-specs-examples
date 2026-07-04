
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;
import com.azure.resourcemanager.network.models.PublicIpAddressDnsSettings;

/**
 * Samples for PublicIpAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressCreateDns.json
     */
    /**
     * Sample code: Create public IP address DNS.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createPublicIPAddressDNS(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().createOrUpdate("rg1", "test-ip", new PublicIpAddressInner()
            .withLocation("eastus").withDnsSettings(new PublicIpAddressDnsSettings().withDomainNameLabel("dnslbl")),
            com.azure.core.util.Context.NONE);
    }
}
