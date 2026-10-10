
import com.azure.resourcemanager.resiliencemanagement.models.ReprotectRequest;
import com.azure.resourcemanager.resiliencemanagement.models.ReprotectRequestProperties;
import java.util.Arrays;

/**
 * Samples for RecoveryPlanActions Reprotect.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_Reprotect_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_Reprotect_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsReprotectMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().reprotect("sampleServiceGroupName", "qmn", "samplePlanName",
            new ReprotectRequest()
                .withReprotectRequestProperties(new ReprotectRequestProperties().withSelectedResourceIds(Arrays.asList(
                    "/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryResources/12345678-9012-3456-7890-123456789012"))),
            com.azure.core.util.Context.NONE);
    }
}
