
/**
 * Samples for WorkloadImpacts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/WorkloadImpact_Get.json
     */
    /**
     * Sample code: Get WorkloadImpact Resource by name example.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void getWorkloadImpactResourceByNameExample(
        com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.workloadImpacts().getWithResponse("impact-001", com.azure.core.util.Context.NONE);
    }
}
