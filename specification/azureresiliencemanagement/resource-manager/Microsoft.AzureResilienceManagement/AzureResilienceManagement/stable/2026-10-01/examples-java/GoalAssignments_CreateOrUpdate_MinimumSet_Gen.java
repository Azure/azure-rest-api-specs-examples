
import com.azure.resourcemanager.resiliencemanagement.fluent.models.GoalAssignmentInner;
import com.azure.resourcemanager.resiliencemanagement.models.GoalAssignmentProperties;

/**
 * Samples for GoalAssignments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalAssignments_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: GoalAssignments_CreateOrUpdate_MinimumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalAssignmentsCreateOrUpdateMinimumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalAssignments().createOrUpdate("production-sg", "zonal-resiliency-goal",
            new GoalAssignmentInner().withProperties(new GoalAssignmentProperties().withRequireZonalResiliency(true)),
            com.azure.core.util.Context.NONE);
    }
}
