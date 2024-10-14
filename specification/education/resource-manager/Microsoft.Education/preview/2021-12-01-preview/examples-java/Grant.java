
/**
 * Samples for Grants Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/Grant.json
     */
    /**
     * Sample code: Grant.
     * 
     * @param manager Entry point to EducationManager.
     */
    public static void grant(com.azure.resourcemanager.education.EducationManager manager) {
        manager.grants().getWithResponse("{billingAccountName}", "{billingProfileName}", false,
            com.azure.core.util.Context.NONE);
    }
}
