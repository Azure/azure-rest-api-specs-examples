
/**
 * Samples for ReplicationProtectedItems Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Get_A2ACvm.json
     */
    /**
     * Sample code: Gets the details of an A2A confidential VM with CMK disk encryption.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfAnA2AConfidentialVMWithCMKDiskEncryption(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().getWithResponse("priyanprg", "priyanponeboxvault", "CentralUSCanSite",
            "CentralUSCancloud", "a2aCvmVm", com.azure.core.util.Context.NONE);
    }
}
