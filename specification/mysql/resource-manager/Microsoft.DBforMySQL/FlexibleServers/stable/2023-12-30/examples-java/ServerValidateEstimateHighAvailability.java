
import com.azure.resourcemanager.mysqlflexibleserver.fluent.models.HighAvailabilityValidationEstimationInner;

/**
 * Samples for Servers ValidateEstimateHighAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * ServerValidateEstimateHighAvailability.json
     */
    /**
     * Sample code: Validate a validation and estimation of high availability.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void validateAValidationAndEstimationOfHighAvailability(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().validateEstimateHighAvailabilityWithResponse("TestGroup", "testserver",
            new HighAvailabilityValidationEstimationInner().withExpectedStandbyAvailabilityZone("1"),
            com.azure.core.util.Context.NONE);
    }
}
