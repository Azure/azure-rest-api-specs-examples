
/**
 * Samples for Enrollments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Enrollments_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Enrollments_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        enrollmentsListMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.enrollments().list("MyResourceGroup", "myUsagePlan", com.azure.core.util.Context.NONE);
    }
}
