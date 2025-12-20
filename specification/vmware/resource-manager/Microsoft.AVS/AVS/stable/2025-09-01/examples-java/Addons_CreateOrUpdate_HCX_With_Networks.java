
import com.azure.resourcemanager.avs.models.AddonHcxProperties;

/**
 * Samples for Addons CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Addons_CreateOrUpdate_HCX_With_Networks.json
     */
    /**
     * Sample code: Addons_CreateOrUpdate_HCX_With_Networks.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsCreateOrUpdateHCXWithNetworks(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().define("hcx").withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new AddonHcxProperties().withOffer("VMware MaaS Cloud Provider (Enterprise)")
                .withManagementNetwork("10.3.1.0/24").withUplinkNetwork("10.3.2.0/24"))
            .create();
    }
}
