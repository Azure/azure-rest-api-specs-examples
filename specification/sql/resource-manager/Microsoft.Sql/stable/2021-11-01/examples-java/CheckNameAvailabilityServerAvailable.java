
import com.azure.resourcemanager.sql.models.CheckNameAvailabilityRequest;

/**
 * Samples for Servers CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CheckNameAvailabilityServerAvailable.
     * json
     */
    /**
     * Sample code: Check for a server name that is available.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkForAServerNameThatIsAvailable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityRequest().withName("server1"), com.azure.core.util.Context.NONE);
    }
}
