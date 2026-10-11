
import com.azure.resourcemanager.resiliencemanagement.models.RecoveryOperationNames;
import com.azure.resourcemanager.resiliencemanagement.models.ValidateForOperationRequest;

/**
 * Samples for RecoveryPlanActions ValidateForOperation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_ValidateForOperation_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_ValidateForOperation_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsValidateForOperationMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().validateForOperation("sampleServiceGroupName", "qmn", "samplePlanName",
            new ValidateForOperationRequest().withOperationName(RecoveryOperationNames.FAILOVER),
            com.azure.core.util.Context.NONE);
    }
}
