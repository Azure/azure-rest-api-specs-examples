
import com.azure.resourcemanager.postgresqlflexibleserver.models.CheckNameAvailabilityRequest;

/**
 * Samples for NameAvailability CheckWithLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/NameAvailabilityCheckWithLocation.json
     */
    /**
     * Sample code: Check the validity and availability of the given name, in the given location, to assign it to a new
     * server or to use it as the base name of a new pair of virtual endpoints.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        checkTheValidityAndAvailabilityOfTheGivenNameInTheGivenLocationToAssignItToANewServerOrToUseItAsTheBaseNameOfANewPairOfVirtualEndpoints(
            com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.nameAvailabilities().checkWithLocationWithResponse("eastus", new CheckNameAvailabilityRequest()
            .withName("exampleserver").withType("Microsoft.DBforPostgreSQL/flexibleServers"),
            com.azure.core.util.Context.NONE);
    }
}
