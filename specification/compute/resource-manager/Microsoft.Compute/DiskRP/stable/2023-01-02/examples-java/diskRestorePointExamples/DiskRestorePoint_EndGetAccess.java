/** Samples for DiskRestorePoint RevokeAccess. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/diskRestorePointExamples/DiskRestorePoint_EndGetAccess.json
     */
    /**
     * Sample code: Revokes access to a diskRestorePoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void revokesAccessToADiskRestorePoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskRestorePoints()
            .revokeAccess(
                "myResourceGroup",
                "rpc",
                "vmrp",
                "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745",
                com.azure.core.util.Context.NONE);
    }
}
