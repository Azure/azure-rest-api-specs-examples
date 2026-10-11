
import com.azure.resourcemanager.resiliencemanagement.models.EnrollmentProperties;

/**
 * Samples for Enrollments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Enrollments_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Enrollments_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void enrollmentsCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.enrollments().define("sg1-enrollment").withExistingUsagePlan("MyResourceGroup", "myUsagePlan")
            .withProperties(
                new EnrollmentProperties().withServiceGroupId("/providers/Microsoft.Management/serviceGroups/sg1"))
            .create();
    }
}
