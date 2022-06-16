import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.OfficeAtpCheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/CheckRequirementsOfficeATP.json
     */
    /**
     * Sample code: Check requirements for OfficeATP.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForOfficeATP(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse("myRg", "myWorkspace", new OfficeAtpCheckRequirements(), Context.NONE);
    }
}
