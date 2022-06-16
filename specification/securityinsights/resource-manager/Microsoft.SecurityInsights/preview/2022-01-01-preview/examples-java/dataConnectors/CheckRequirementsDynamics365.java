import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.Dynamics365CheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/CheckRequirementsDynamics365.json
     */
    /**
     * Sample code: Check requirements for Dynamics365.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForDynamics365(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse("myRg", "myWorkspace", new Dynamics365CheckRequirements(), Context.NONE);
    }
}
