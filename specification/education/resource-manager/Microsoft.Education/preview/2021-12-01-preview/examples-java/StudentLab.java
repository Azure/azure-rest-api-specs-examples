import com.azure.core.util.Context;

/** Samples for StudentLabs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentLab.json
     */
    /**
     * Sample code: StudentLab.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void studentLab(com.azure.resourcemanager.education.EducationManager manager) {
        manager.studentLabs().getWithResponse("{studentLabName}", Context.NONE);
    }
}
