import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.OfficePowerBICheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CheckRequirementsOfficePowerBI.json
     */
    /**
     * Sample code: Check requirements for OfficePowerBI.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForOfficePowerBI(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse("myRg", "myWorkspace", new OfficePowerBICheckRequirements(), Context.NONE);
    }
}
