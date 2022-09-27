import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.McasCheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CheckRequirementsMdatp.json
     */
    /**
     * Sample code: Check requirements for Mdatp.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForMdatp(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse("myRg", "myWorkspace", new McasCheckRequirements(), Context.NONE);
    }
}
