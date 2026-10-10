
import com.azure.resourcemanager.resiliencemanagement.models.RecommendCapacityRequest;
import java.util.Arrays;

/**
 * Samples for GoalAssignments RecommendCapacity.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalAssignments_RecommendCapacity_MaximumSet_Gen.json
     */
    /**
     * Sample code: GoalAssignments_RecommendCapacity_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalAssignmentsRecommendCapacityMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalAssignments().recommendCapacity("production-sg", "zonal-resiliency-goal",
            new RecommendCapacityRequest().withResourceIds(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/vm1",
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRg/providers/Microsoft.Storage/storageAccounts/sa1")),
            com.azure.core.util.Context.NONE);
    }
}
