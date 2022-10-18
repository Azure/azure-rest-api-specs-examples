import com.azure.core.util.Context;

/** Samples for HcxEnterpriseSites Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/HcxEnterpriseSites_Delete.json
     */
    /**
     * Sample code: HcxEnterpriseSites_Delete.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void hcxEnterpriseSitesDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.hcxEnterpriseSites().deleteWithResponse("group1", "cloud1", "site1", Context.NONE);
    }
}
