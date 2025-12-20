
import com.azure.resourcemanager.avs.models.AddonHcxProperties;

/**
 * Samples for Addons CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Addons_CreateOrUpdate_HCX.json
     */
    /**
     * Sample code: Addons_CreateOrUpdate_HCX.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsCreateOrUpdateHCX(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().define("hcx").withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new AddonHcxProperties().withOffer("VMware MaaS Cloud Provider (Enterprise)")).create();
    }
}
