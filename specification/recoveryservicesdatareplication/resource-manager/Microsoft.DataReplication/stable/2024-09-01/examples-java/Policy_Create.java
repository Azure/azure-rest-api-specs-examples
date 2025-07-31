
/**
 * Samples for Policy Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Policy_Create.json
     */
    /**
     * Sample code: Puts the policy.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void putsThePolicy(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.policies().define("fafqwc").withExistingReplicationVault("rgrecoveryservicesdatareplication", "4")
            .create();
    }
}
