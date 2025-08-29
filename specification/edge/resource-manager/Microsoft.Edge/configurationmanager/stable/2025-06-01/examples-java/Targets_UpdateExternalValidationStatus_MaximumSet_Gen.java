
import com.azure.core.management.exception.ManagementError;
import com.azure.resourcemanager.workloadorchestration.models.UpdateExternalValidationStatusParameter;
import com.azure.resourcemanager.workloadorchestration.models.ValidationStatus;

/**
 * Samples for Targets UpdateExternalValidationStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_UpdateExternalValidationStatus_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_UpdateExternalValidationStatus_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void targetsUpdateExternalValidationStatusMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().updateExternalValidationStatus("rgconfigurationmanager", "testname",
            new UpdateExternalValidationStatusParameter().withSolutionVersionId("shntcsuwlmpehmuqkrbf")
                .withErrorDetails(new ManagementError()).withExternalValidationId("ivsjzwy")
                .withValidationStatus(ValidationStatus.VALID),
            com.azure.core.util.Context.NONE);
    }
}
