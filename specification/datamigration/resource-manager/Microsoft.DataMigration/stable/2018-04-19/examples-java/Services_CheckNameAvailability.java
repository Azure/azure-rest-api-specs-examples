import com.azure.resourcemanager.datamigration.models.NameAvailabilityRequest;

/** Samples for Services CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Services_CheckNameAvailability.json
     */
    /**
     * Sample code: Services_CheckNameAvailability.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesCheckNameAvailability(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager
            .services()
            .checkNameAvailabilityWithResponse(
                "eastus", new NameAvailabilityRequest(), com.azure.core.util.Context.NONE);
    }
}
