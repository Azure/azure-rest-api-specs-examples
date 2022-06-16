import com.azure.core.util.Context;

/** Samples for Students List. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentList.json
     */
    /**
     * Sample code: StudentList.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void studentList(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .students()
            .list("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", null, Context.NONE);
    }
}
