
/**
 * Samples for TieringCostOperationStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * TieringCost/GetTieringCostOperationStatus.json
     */
    /**
     * Sample code: Fetch Tiering Cost Operation Status.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void fetchTieringCostOperationStatus(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.tieringCostOperationStatus().getWithResponse("gaallaRG", "gaallavaultbvtd2msi",
            "0f48183b-0a44-4dca-aec1-bba5daab888a", com.azure.core.util.Context.NONE);
    }
}
