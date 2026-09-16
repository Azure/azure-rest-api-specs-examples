
import com.azure.resourcemanager.securityinsights.models.MSTICheckRequirements;

/**
 * Samples for DataConnectorsCheckRequirementsOperation Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CheckRequirementsMicrosoftThreatIntelligence.json
     */
    /**
     * Sample code: Check requirements for MicrosoftThreatIntelligence.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void checkRequirementsForMicrosoftThreatIntelligence(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorsCheckRequirementsOperations().postWithResponse("myRg", "myWorkspace",
            new MSTICheckRequirements().withTenantId("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
            com.azure.core.util.Context.NONE);
    }
}
