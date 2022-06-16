import com.azure.core.util.Context;

/** Samples for Grants List. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GrantDefaultList.json
     */
    /**
     * Sample code: GrantList.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void grantList(com.azure.resourcemanager.education.EducationManager manager) {
        manager.grants().list("{billingAccountName}", "{billingProfileName}", false, Context.NONE);
    }
}
