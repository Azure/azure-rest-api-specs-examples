
import com.azure.resourcemanager.resiliencemanagement.models.FailoverDirectionTypes;
import com.azure.resourcemanager.resiliencemanagement.models.FailoverRequest;
import com.azure.resourcemanager.resiliencemanagement.models.FailoverRequestProperties;
import java.util.Arrays;

/**
 * Samples for RecoveryPlanActions ValidateForTestFailover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_ValidateForTestFailover_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_ValidateForTestFailover_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsValidateForTestFailoverMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().validateForTestFailover("sampleServiceGroupName", "qmn", "samplePlanName",
            new FailoverRequest().withFailoverDirection(FailoverDirectionTypes.FROM_SPECIFIC_LOCATIONS)
                .withFailoverRequestProperties(
                    new FailoverRequestProperties().withSourceLocations(Arrays.asList("westus"))),
            com.azure.core.util.Context.NONE);
    }
}
