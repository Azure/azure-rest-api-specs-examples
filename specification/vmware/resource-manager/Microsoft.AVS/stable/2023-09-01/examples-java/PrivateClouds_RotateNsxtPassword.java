
/**
 * Samples for PrivateClouds RotateNsxtPassword.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/PrivateClouds_RotateNsxtPassword.
     * json
     */
    /**
     * Sample code: PrivateClouds_RotateNsxtPassword.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsRotateNsxtPassword(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().rotateNsxtPassword("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
