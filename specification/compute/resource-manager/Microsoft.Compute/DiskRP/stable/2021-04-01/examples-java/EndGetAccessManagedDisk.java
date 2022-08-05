import com.azure.core.util.Context;

/** Samples for Disks RevokeAccess. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/EndGetAccessManagedDisk.json
     */
    /**
     * Sample code: Revoke access to a managed disk.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void revokeAccessToAManagedDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .revokeAccess("myResourceGroup", "myDisk", Context.NONE);
    }
}
