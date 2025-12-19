
/**
 * Samples for HcxEnterpriseSites Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/HcxEnterpriseSites_Delete.json
     */
    /**
     * Sample code: HcxEnterpriseSites_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void hcxEnterpriseSitesDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.hcxEnterpriseSites().deleteWithResponse("group1", "cloud1", "site1", com.azure.core.util.Context.NONE);
    }
}
