
/**
 * Samples for Enrollments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Enrollments_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Enrollments_Delete_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void enrollmentsDeleteMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.enrollments().delete("MyResourceGroup", "myUsagePlan", "sg1-enrollment",
            com.azure.core.util.Context.NONE);
    }
}
