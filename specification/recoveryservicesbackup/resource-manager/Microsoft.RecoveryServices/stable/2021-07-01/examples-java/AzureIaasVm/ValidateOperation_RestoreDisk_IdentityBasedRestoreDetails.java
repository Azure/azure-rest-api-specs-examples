import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.ValidateOperationRequest;

/** Samples for OperationOperation Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureIaasVm/ValidateOperation_RestoreDisk_IdentityBasedRestoreDetails.json
     */
    /**
     * Sample code: Validate Operation with identityBasedRestoreDetails.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void validateOperationWithIdentityBasedRestoreDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .operationOperations()
            .validateWithResponse("testVault", "testRG", new ValidateOperationRequest(), Context.NONE);
    }
}
