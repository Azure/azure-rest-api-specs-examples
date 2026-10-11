
/**
 * Samples for Enrollments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Enrollments_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Enrollments_Get_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        enrollmentsGetMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.enrollments().getWithResponse("MyResourceGroup", "myUsagePlan", "sg1-enrollment",
            com.azure.core.util.Context.NONE);
    }
}
