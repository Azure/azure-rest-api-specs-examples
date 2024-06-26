
/**
 * Samples for PrivateClouds GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/PrivateClouds_Get_Stretched.json
     */
    /**
     * Sample code: PrivateClouds_Get_Stretched.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsGetStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().getByResourceGroupWithResponse("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
