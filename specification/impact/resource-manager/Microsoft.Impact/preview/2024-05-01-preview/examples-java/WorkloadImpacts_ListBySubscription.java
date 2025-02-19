
/**
 * Samples for WorkloadImpacts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/WorkloadImpacts_ListBySubscription.json
     */
    /**
     * Sample code: List WorkloadImpact resources by subscription.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void listWorkloadImpactResourcesBySubscription(
        com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.workloadImpacts().list(com.azure.core.util.Context.NONE);
    }
}
