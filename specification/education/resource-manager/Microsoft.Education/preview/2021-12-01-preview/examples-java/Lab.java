import com.azure.core.util.Context;

/** Samples for Labs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/Lab.json
     */
    /**
     * Sample code: Lab.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void lab(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .labs()
            .getWithResponse(
                "{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", false, Context.NONE);
    }
}
