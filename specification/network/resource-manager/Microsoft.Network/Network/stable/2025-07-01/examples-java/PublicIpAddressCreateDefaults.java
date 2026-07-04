
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;

/**
 * Samples for PublicIpAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressCreateDefaults.json
     */
    /**
     * Sample code: Create public IP address defaults.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createPublicIPAddressDefaults(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().createOrUpdate("rg1", "test-ip",
            new PublicIpAddressInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}
