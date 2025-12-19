
import com.azure.resourcemanager.avs.models.LicenseName;

/**
 * Samples for Licenses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Licenses_CreateOrUpdate.json
     */
    /**
     * Sample code: Licenses_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void licensesCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.licenses().define(LicenseName.VMWARE_FIREWALL).withExistingPrivateCloud("group1", "cloud1").create();
    }
}
