import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.AadCheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CheckRequirementsAzureActiveDirectoryNoAuthorization.json
     */
    /**
     * Sample code: Check requirements for AAD - no authorization.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForAADNoAuthorization(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse("myRg", "myWorkspace", new AadCheckRequirements(), Context.NONE);
    }
}
