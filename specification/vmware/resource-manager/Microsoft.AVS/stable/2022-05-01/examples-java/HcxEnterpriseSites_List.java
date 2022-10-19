import com.azure.core.util.Context;

/** Samples for HcxEnterpriseSites List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/HcxEnterpriseSites_List.json
     */
    /**
     * Sample code: HcxEnterpriseSites_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void hcxEnterpriseSitesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.hcxEnterpriseSites().list("group1", "cloud1", Context.NONE);
    }
}
