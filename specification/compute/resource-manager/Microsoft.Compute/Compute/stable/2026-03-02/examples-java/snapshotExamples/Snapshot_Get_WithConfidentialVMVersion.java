
/**
 * Samples for Snapshots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/snapshotExamples/Snapshot_Get_WithConfidentialVMVersion.json
     */
    /**
     * Sample code: Get information about a ConfidentialVM snapshot with confidentialVMVersion.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutAConfidentialVMSnapshotWithConfidentialVMVersion(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().getByResourceGroupWithResponse("myResourceGroup",
            "myConfidentialSnapshot", com.azure.core.util.Context.NONE);
    }
}
