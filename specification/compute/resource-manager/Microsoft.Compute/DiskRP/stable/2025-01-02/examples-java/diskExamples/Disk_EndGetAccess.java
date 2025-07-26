
/**
 * Samples for Disks RevokeAccess.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskExamples/
     * Disk_EndGetAccess.json
     */
    /**
     * Sample code: revoke access to a managed disk.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void revokeAccessToAManagedDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().revokeAccess("myResourceGroup", "myDisk",
            com.azure.core.util.Context.NONE);
    }
}
