
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Operations_List.json
     */
    /**
     * Sample code: List operations for Microsoft.ProgramEnrollment.
     * 
     * @param manager Entry point to ProgramEnrollmentManager.
     */
    public static void listOperationsForMicrosoftProgramEnrollment(
        com.azure.resourcemanager.programenrollment.ProgramEnrollmentManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
