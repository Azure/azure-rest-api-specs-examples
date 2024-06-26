
/**
 * Samples for PrivateClouds RotateVcenterPassword.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/
     * PrivateClouds_RotateVcenterPassword.json
     */
    /**
     * Sample code: PrivateClouds_RotateVcenterPassword.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsRotateVcenterPassword(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().rotateVcenterPassword("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
