
/**
 * Samples for EduEnrollments GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EduEnrollments_Get.json
     */
    /**
     * Sample code: Get an edu enrollment.
     * 
     * @param manager Entry point to ProgramEnrollmentManager.
     */
    public static void
        getAnEduEnrollment(com.azure.resourcemanager.programenrollment.ProgramEnrollmentManager manager) {
        manager.eduEnrollments().getByResourceGroupWithResponse("testrg", "default", com.azure.core.util.Context.NONE);
    }
}
