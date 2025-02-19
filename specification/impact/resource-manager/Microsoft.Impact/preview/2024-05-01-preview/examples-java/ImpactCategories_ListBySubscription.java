
/**
 * Samples for ImpactCategories List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/ImpactCategories_ListBySubscription.json
     */
    /**
     * Sample code: Get ImpactCategories list by subscription.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void getImpactCategoriesListBySubscription(
        com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.impactCategories().list("microsoft.compute/virtualmachines", null, com.azure.core.util.Context.NONE);
    }
}
