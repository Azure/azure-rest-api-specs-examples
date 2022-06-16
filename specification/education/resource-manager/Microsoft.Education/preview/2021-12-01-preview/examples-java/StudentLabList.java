import com.azure.core.util.Context;

/** Samples for StudentLabs ListAll. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentLabList.json
     */
    /**
     * Sample code: StudentLabList.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void studentLabList(com.azure.resourcemanager.education.EducationManager manager) {
        manager.studentLabs().listAll(Context.NONE);
    }
}
