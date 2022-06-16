import com.azure.core.util.Context;

/** Samples for Students Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/Student.json
     */
    /**
     * Sample code: Student.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void student(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .students()
            .getWithResponse(
                "{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", "{studentAlias}", Context.NONE);
    }
}
