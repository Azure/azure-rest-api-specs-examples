/** Samples for HcxEnterpriseSites CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/HcxEnterpriseSites_CreateOrUpdate.json
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
