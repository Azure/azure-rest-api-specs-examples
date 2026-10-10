
import com.azure.resourcemanager.resiliencemanagement.fluent.models.GoalAssignmentInner;
import com.azure.resourcemanager.resiliencemanagement.models.GoalAssignmentProperties;
import com.azure.resourcemanager.resiliencemanagement.models.ServiceLevelResource;
import java.util.Arrays;

/**
 * Samples for GoalAssignments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalAssignments_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: GoalAssignments_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalAssignmentsCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalAssignments().createOrUpdate("production-sg", "zonal-resiliency-goal",
            new GoalAssignmentInner().withProperties(new GoalAssignmentProperties().withRequireZonalResiliency(true)
                .withServiceLevelResources(Arrays.asList(new ServiceLevelResource().withServiceLevelIndicatorResourceId(
                    "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.Compute/virtualMachines/MyVirtualMachine")))),
            com.azure.core.util.Context.NONE);
    }
}
