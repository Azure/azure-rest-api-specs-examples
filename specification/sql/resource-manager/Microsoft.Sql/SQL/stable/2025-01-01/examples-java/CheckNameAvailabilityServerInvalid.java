
import com.azure.resourcemanager.sql.models.CheckNameAvailabilityRequest;

/**
 * Samples for Servers CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CheckNameAvailabilityServerInvalid.json
     */
    /**
     * Sample code: Check for a server name that is invalid.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void checkForAServerNameThatIsInvalid(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityRequest().withName("SERVER1"), com.azure.core.util.Context.NONE);
    }
}
