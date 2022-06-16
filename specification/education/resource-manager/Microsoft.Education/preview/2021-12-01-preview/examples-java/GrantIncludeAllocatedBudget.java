import com.azure.core.util.Context;

/** Samples for Grants Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GrantIncludeAllocatedBudget.json
     */
    /**
     * Sample code: GrantIncludeAllocatedBudget.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void grantIncludeAllocatedBudget(com.azure.resourcemanager.education.EducationManager manager) {
        manager.grants().getWithResponse("{billingAccountName}", "{billingProfileName}", false, Context.NONE);
    }
}
