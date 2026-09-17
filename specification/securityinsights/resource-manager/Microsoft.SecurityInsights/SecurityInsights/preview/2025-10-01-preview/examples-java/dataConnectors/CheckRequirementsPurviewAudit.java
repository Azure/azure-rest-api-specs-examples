
import com.azure.resourcemanager.securityinsights.models.PurviewAuditCheckRequirements;

/**
 * Samples for DataConnectorsCheckRequirementsOperation Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CheckRequirementsPurviewAudit.json
     */
    /**
     * Sample code: Check requirements for PurviewAudit.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        checkRequirementsForPurviewAudit(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorsCheckRequirementsOperations().postWithResponse("myRg", "myWorkspace",
            new PurviewAuditCheckRequirements().withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
            com.azure.core.util.Context.NONE);
    }
}
