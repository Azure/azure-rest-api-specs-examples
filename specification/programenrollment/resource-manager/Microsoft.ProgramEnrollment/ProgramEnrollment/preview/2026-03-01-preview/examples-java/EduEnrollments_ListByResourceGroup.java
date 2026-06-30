
/**
 * Samples for EduEnrollments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EduEnrollments_ListByResourceGroup.json
     */
    /**
     * Sample code: List edu enrollments by resource group.
     * 
     * @param manager Entry point to ProgramEnrollmentManager.
     */
    public static void listEduEnrollmentsByResourceGroup(
        com.azure.resourcemanager.programenrollment.ProgramEnrollmentManager manager) {
        manager.eduEnrollments().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
