import com.azure.core.util.Context;

/** Samples for Labs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/LabIncludeBudget.json
     */
    /**
     * Sample code: LabIncludeBudget.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void labIncludeBudget(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .labs()
            .getWithResponse(
                "{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", true, Context.NONE);
    }
}
