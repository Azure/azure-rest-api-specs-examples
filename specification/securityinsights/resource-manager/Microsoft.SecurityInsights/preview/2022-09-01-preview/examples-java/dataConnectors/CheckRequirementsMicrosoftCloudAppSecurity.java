
import com.azure.resourcemanager.securityinsights.models.McasCheckRequirements;

/**
 * Samples for DataConnectorsCheckRequirementsOperation Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * dataConnectors/CheckRequirementsMicrosoftCloudAppSecurity.json
     */
    /**
     * Sample code: Check requirements for Mcas.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        checkRequirementsForMcas(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorsCheckRequirementsOperations().postWithResponse("myRg", "myWorkspace",
            new McasCheckRequirements().withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
            com.azure.core.util.Context.NONE);
    }
}
