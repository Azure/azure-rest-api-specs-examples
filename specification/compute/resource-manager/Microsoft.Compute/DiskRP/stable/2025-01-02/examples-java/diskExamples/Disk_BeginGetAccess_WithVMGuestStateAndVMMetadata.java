
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/**
 * Samples for Disks GrantAccess.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskExamples/
     * Disk_BeginGetAccess_WithVMGuestStateAndVMMetadata.json
     */
    /**
     * Sample code: get sas on managed disk, VM guest state and VM metadata for Confidential VM.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSasOnManagedDiskVMGuestStateAndVMMetadataForConfidentialVM(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().grantAccess("myResourceGroup", "myDisk",
            new GrantAccessData().withAccess(AccessLevel.READ).withDurationInSeconds(300)
                .withGetSecureVMGuestStateSas(true),
            com.azure.core.util.Context.NONE);
    }
}
