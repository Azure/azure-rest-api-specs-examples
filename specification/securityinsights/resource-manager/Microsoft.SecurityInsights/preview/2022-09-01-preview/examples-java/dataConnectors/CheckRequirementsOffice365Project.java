import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.Office365ProjectCheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CheckRequirementsOffice365Project.json
     */
    /**
     * Sample code: Check requirements for Office365Project.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForOffice365Project(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse("myRg", "myWorkspace", new Office365ProjectCheckRequirements(), Context.NONE);
    }
}
