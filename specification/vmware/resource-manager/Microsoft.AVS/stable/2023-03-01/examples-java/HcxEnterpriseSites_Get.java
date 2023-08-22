/** Samples for HcxEnterpriseSites Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/HcxEnterpriseSites_Get.json
     */
    /**
     * Sample code: HcxEnterpriseSites_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void hcxEnterpriseSitesGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.hcxEnterpriseSites().getWithResponse("group1", "cloud1", "site1", com.azure.core.util.Context.NONE);
    }
}
