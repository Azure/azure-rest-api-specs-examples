
/**
 * Samples for EduEnrollments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EduEnrollments_ListBySubscription.json
     */
    /**
     * Sample code: List edu enrollments by subscription.
     * 
     * @param manager Entry point to ProgramEnrollmentManager.
     */
    public static void
        listEduEnrollmentsBySubscription(com.azure.resourcemanager.programenrollment.ProgramEnrollmentManager manager) {
        manager.eduEnrollments().list(com.azure.core.util.Context.NONE);
    }
}
