
/**
 * Samples for Disks RevokeAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_EndGetAccess.json
     */
    /**
     * Sample code: revoke access to a managed disk.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void revokeAccessToAManagedDisk(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().revokeAccess("myResourceGroup", "myDisk", com.azure.core.util.Context.NONE);
    }
}
