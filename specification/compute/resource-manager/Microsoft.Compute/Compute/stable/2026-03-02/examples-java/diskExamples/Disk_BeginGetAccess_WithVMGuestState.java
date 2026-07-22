
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/**
 * Samples for Disks GrantAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_BeginGetAccess_WithVMGuestState.json
     */
    /**
     * Sample code: get sas on managed disk and VM guest state.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getSasOnManagedDiskAndVMGuestState(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().grantAccess("myResourceGroup", "myDisk", new GrantAccessData()
            .withAccess(AccessLevel.READ).withDurationInSeconds(300).withGetSecureVMGuestStateSas(true),
            com.azure.core.util.Context.NONE);
    }
}
