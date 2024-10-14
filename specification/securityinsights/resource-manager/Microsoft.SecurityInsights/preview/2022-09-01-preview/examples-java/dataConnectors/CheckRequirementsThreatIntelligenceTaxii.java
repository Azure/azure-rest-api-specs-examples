
import com.azure.resourcemanager.securityinsights.models.TiTaxiiCheckRequirements;

/**
 * Samples for DataConnectorsCheckRequirementsOperation Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * dataConnectors/CheckRequirementsThreatIntelligenceTaxii.json
     */
    /**
     * Sample code: Check requirements for TI Taxii.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        checkRequirementsForTITaxii(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorsCheckRequirementsOperations().postWithResponse("myRg", "myWorkspace",
            new TiTaxiiCheckRequirements().withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
            com.azure.core.util.Context.NONE);
    }
}
