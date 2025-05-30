
/**
 * Samples for GetTieringCostOperationResult Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/
     * TieringCost/GetTieringCostOperationResult.json
     */
    /**
     * Sample code: Fetch Tiering Cost Operation Result.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void fetchTieringCostOperationResult(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.getTieringCostOperationResults().getWithResponse("gaallaRG", "gaallavaultbvtd2msi",
            "0f48183b-0a44-4dca-aec1-bba5daab888a", com.azure.core.util.Context.NONE);
    }
}
