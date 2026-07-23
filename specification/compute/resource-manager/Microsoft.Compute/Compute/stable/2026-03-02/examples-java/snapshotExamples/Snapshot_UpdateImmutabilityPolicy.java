
import com.azure.resourcemanager.compute.models.ImmutabilityPolicyData;
import com.azure.resourcemanager.compute.models.ImmutabilityPolicyType;

/**
 * Samples for Snapshots UpdateImmutabilityPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/snapshotExamples/Snapshot_UpdateImmutabilityPolicy.json
     */
    /**
     * Sample code: Update immutability policy of a snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateImmutabilityPolicyOfASnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().updateImmutabilityPolicy("myResourceGroup", "mySnapshot",
            new ImmutabilityPolicyData().withImmutabilityDurationDays(30).withType(ImmutabilityPolicyType.UNLOCKED),
            com.azure.core.util.Context.NONE);
    }
}
