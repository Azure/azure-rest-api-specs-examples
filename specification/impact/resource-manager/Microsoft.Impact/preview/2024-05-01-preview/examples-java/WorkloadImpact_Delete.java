
/**
 * Samples for WorkloadImpacts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/WorkloadImpact_Delete.json
     */
    /**
     * Sample code: Delete WorkloadImpact Resource by name example.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void deleteWorkloadImpactResourceByNameExample(
        com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.workloadImpacts().deleteWithResponse("impact-001", com.azure.core.util.Context.NONE);
    }
}
