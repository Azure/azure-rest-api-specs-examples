import com.azure.core.util.Context;

/** Samples for Grants ListAll. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GrantListIncludeAllocatedBudget.json
     */
    /**
     * Sample code: GrantListIncludeAllocatedBudget.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void grantListIncludeAllocatedBudget(com.azure.resourcemanager.education.EducationManager manager) {
        manager.grants().listAll(true, Context.NONE);
    }
}
