import com.azure.resourcemanager.datamigration.models.NameAvailabilityRequest;

/** Samples for Services NestedCheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Services_NestedCheckNameAvailability.json
     */
    /**
     * Sample code: Services_NestedCheckNameAvailability.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesNestedCheckNameAvailability(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager
            .services()
            .nestedCheckNameAvailabilityWithResponse(
                "DmsSdkRg", "DmsSdkService", new NameAvailabilityRequest(), com.azure.core.util.Context.NONE);
    }
}
