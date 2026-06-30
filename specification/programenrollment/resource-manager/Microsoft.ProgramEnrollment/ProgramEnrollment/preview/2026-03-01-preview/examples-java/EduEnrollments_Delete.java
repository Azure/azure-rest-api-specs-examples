
/**
 * Samples for EduEnrollments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EduEnrollments_Delete.json
     */
    /**
     * Sample code: Delete an edu enrollment.
     * 
     * @param manager Entry point to ProgramEnrollmentManager.
     */
    public static void
        deleteAnEduEnrollment(com.azure.resourcemanager.programenrollment.ProgramEnrollmentManager manager) {
        manager.eduEnrollments().delete("testrg", "default", com.azure.core.util.Context.NONE);
    }
}
