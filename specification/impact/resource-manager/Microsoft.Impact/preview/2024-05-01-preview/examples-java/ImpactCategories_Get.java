
/**
 * Samples for ImpactCategories Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/ImpactCategories_Get.json
     */
    /**
     * Sample code: Get WorkloadImpact Resource by name.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void
        getWorkloadImpactResourceByName(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.impactCategories().getWithResponse("ARMOperation.Create", com.azure.core.util.Context.NONE);
    }
}
