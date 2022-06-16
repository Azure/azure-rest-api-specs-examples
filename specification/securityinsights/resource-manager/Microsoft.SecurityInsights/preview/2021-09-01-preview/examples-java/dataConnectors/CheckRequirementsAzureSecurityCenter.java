import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.AscCheckRequirements;

/** Samples for DataConnectorsCheckRequirementsOperation Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/CheckRequirementsAzureSecurityCenter.json
     */
    /**
     * Sample code: Check requirements for ASC.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForASC(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectorsCheckRequirementsOperations()
            .postWithResponse(
                "myRg",
                "myWorkspace",
                new AscCheckRequirements().withSubscriptionId("c0688291-89d7-4bed-87a2-a7b1bff43f4c"),
                Context.NONE);
    }
}
