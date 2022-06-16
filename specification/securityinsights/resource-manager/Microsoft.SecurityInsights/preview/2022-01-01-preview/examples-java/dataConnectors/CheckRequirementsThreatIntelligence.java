import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.TICheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/CheckRequirementsThreatIntelligence.json
     */
    /**
     * Sample code: Check requirements for TI.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForTI(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse("myRg", "myWorkspace", new TICheckRequirements(), Context.NONE);
    }
}
