
import com.azure.resourcemanager.datamigration.models.NameAvailabilityRequest;

/**
 * Samples for Services CheckChildrenNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * Services_CheckChildrenNameAvailability.json
     */
    /**
     * Sample code: Services_CheckChildrenNameAvailability.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void
        servicesCheckChildrenNameAvailability(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.services().checkChildrenNameAvailabilityWithResponse("DmsSdkRg", "DmsSdkService",
            new NameAvailabilityRequest().withName("Task1").withType("tasks"), com.azure.core.util.Context.NONE);
    }
}
