
import com.azure.resourcemanager.postgresqlflexibleserver.models.CheckNameAvailabilityRequest;

/**
 * Samples for NameAvailability CheckGlobally.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * NameAvailabilityCheckGlobally.json
     */
    /**
     * Sample code: Check the validity and availability of the given name, to assign it to a new server or to use it as
     * the base name of a new pair of virtual endpoints.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        checkTheValidityAndAvailabilityOfTheGivenNameToAssignItToANewServerOrToUseItAsTheBaseNameOfANewPairOfVirtualEndpoints(
            com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.nameAvailabilities().checkGloballyWithResponse(new CheckNameAvailabilityRequest()
            .withName("exampleserver").withType("Microsoft.DBforPostgreSQL/flexibleServers"),
            com.azure.core.util.Context.NONE);
    }
}
