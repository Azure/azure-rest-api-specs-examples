
/**
 * Samples for HcxEnterpriseSites CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/HcxEnterpriseSites_CreateOrUpdate.json
     */
    /**
     * Sample code: HcxEnterpriseSites_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void hcxEnterpriseSitesCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.hcxEnterpriseSites().define("site1").withExistingPrivateCloud("group1", "cloud1").create();
    }
}
