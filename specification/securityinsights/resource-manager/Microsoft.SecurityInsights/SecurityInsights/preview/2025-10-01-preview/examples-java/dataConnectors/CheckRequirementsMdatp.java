
import com.azure.resourcemanager.securityinsights.models.MCASCheckRequirements;

/**
 * Samples for DataConnectorsCheckRequirementsOperation Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CheckRequirementsMdatp.json
     */
    /**
     * Sample code: Check requirements for Mdatp.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        checkRequirementsForMdatp(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorsCheckRequirementsOperations().postWithResponse("myRg", "myWorkspace",
            new MCASCheckRequirements().withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
            com.azure.core.util.Context.NONE);
    }
}
