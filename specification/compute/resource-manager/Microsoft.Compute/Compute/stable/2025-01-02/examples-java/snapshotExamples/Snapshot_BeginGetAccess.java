
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.FileFormat;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/**
 * Samples for Snapshots GrantAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_BeginGetAccess.json
     */
    /**
     * Sample code: Get a sas on a snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getASasOnASnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().grantAccess("myResourceGroup", "mySnapshot", new GrantAccessData()
            .withAccess(AccessLevel.READ).withDurationInSeconds(300).withFileFormat(FileFormat.VHDX),
            com.azure.core.util.Context.NONE);
    }
}
