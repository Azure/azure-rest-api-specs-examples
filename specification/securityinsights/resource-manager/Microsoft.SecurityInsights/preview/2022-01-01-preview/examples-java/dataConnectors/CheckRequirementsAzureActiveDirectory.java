import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.AadCheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/CheckRequirementsAzureActiveDirectory.json
     */
    /**
     * Sample code: Check requirements for AAD.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForAAD(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse("myRg", "myWorkspace", new AadCheckRequirements(), Context.NONE);
    }
}
