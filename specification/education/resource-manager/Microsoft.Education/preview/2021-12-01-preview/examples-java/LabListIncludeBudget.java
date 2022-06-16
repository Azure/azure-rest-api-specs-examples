import com.azure.core.util.Context;

/** Samples for Labs ListAll. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/LabListIncludeBudget.json
     */
    /**
     * Sample code: LabListIncludeBudget.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void labListIncludeBudget(com.azure.resourcemanager.education.EducationManager manager) {
        manager.labs().listAll("{billingAccountName}", "{billingProfileName}", true, null, Context.NONE);
    }
}
