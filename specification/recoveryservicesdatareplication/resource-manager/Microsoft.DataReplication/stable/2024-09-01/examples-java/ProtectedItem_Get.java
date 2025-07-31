
/**
 * Samples for ProtectedItem Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProtectedItem_Get.json
     */
    /**
     * Sample code: Gets the protected item.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsTheProtectedItem(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.protectedItems().getWithResponse("rgrecoveryservicesdatareplication", "4", "d",
            com.azure.core.util.Context.NONE);
    }
}
