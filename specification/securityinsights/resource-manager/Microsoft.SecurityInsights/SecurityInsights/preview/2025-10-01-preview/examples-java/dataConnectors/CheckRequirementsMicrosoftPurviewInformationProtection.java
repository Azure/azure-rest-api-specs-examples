
import com.azure.resourcemanager.securityinsights.models.MicrosoftPurviewInformationProtectionCheckRequirements;

/**
 * Samples for DataConnectorsCheckRequirementsOperation Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CheckRequirementsMicrosoftPurviewInformationProtection.json
     */
    /**
     * Sample code: Check requirements for MicrosoftPurviewInformationProtection.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForMicrosoftPurviewInformationProtection(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorsCheckRequirementsOperations().postWithResponse("myRg", "myWorkspace",
            new MicrosoftPurviewInformationProtectionCheckRequirements()
                .withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
            com.azure.core.util.Context.NONE);
    }
}
