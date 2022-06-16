import com.azure.core.util.Context;

/** Samples for Labs ListAll. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/LabList.json
     */
    /**
     * Sample code: LabList.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void labList(com.azure.resourcemanager.education.EducationManager manager) {
        manager.labs().listAll("{billingAccountName}", "{billingProfileName}", false, null, Context.NONE);
    }
}
