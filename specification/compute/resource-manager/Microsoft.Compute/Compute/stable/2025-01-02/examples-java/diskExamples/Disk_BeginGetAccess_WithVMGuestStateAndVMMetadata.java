
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/**
 * Samples for Disks GrantAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_BeginGetAccess_WithVMGuestStateAndVMMetadata.json
     */
    /**
     * Sample code: get sas on managed disk, VM guest state and VM metadata for Confidential VM.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getSasOnManagedDiskVMGuestStateAndVMMetadataForConfidentialVM(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().grantAccess("myResourceGroup", "myDisk", new GrantAccessData()
            .withAccess(AccessLevel.READ).withDurationInSeconds(300).withGetSecureVMGuestStateSas(true),
            com.azure.core.util.Context.NONE);
    }
}
