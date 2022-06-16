import com.azure.core.util.Context;

/** Samples for Labs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/DeleteLab.json
     */
    /**
     * Sample code: DeleteLab.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void deleteLab(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .labs()
            .deleteWithResponse("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", Context.NONE);
    }
}
