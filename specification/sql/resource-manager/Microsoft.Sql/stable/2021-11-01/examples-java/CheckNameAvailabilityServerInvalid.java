
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.CheckNameAvailabilityRequest;

/** Samples for Servers CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CheckNameAvailabilityServerInvalid.
     * json
     */
    /**
     * Sample code: Check for a server name that is invalid.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkForAServerNameThatIsInvalid(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers()
            .checkNameAvailabilityWithResponse(new CheckNameAvailabilityRequest().withName("SERVER1"), Context.NONE);
    }
}
