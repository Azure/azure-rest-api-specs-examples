
import com.azure.resourcemanager.compute.models.ImmutabilityPolicyLockData;
import com.azure.resourcemanager.compute.models.ImmutabilityPolicyType;

/**
 * Samples for Snapshots UpdateImmutabilityPolicyLock.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/snapshotExamples/Snapshot_UpdateImmutabilityPolicyLock.json
     */
    /**
     * Sample code: Update immutability policy lock of a snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        updateImmutabilityPolicyLockOfASnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().updateImmutabilityPolicyLock("myResourceGroup", "mySnapshot",
            new ImmutabilityPolicyLockData().withImmutabilityDurationDays(30).withType(ImmutabilityPolicyType.LOCKED),
            com.azure.core.util.Context.NONE);
    }
}
