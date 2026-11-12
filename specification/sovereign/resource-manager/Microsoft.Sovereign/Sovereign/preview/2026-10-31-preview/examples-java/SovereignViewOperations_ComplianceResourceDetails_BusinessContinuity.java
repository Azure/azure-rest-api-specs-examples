
import com.azure.resourcemanager.sovereign.models.ComplianceDimension;
import com.azure.resourcemanager.sovereign.models.SovereignViewComplianceResourceDetailsContent;
import java.time.OffsetDateTime;

/**
 * Samples for SovereignViewOperations ComplianceResourceDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-31-preview/SovereignViewOperations_ComplianceResourceDetails_BusinessContinuity.json
     */
    /**
     * Sample code: SovereignViewOperations_ComplianceResourceDetails_BusinessContinuity.
     * 
     * @param manager Entry point to SovereignManager.
     */
    public static void sovereignViewOperationsComplianceResourceDetailsBusinessContinuity(
        com.azure.resourcemanager.sovereign.SovereignManager manager) {
        manager.sovereignViewOperations().complianceResourceDetailsWithResponse("ExampleResourceGroup", "ExampleSV",
            new SovereignViewComplianceResourceDetailsContent().withResourceId(
                "/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Sql/servers/exampleserver/databases/exampledatabase")
                .withDimension(ComplianceDimension.BUSINESS_CONTINUITY)
                .withEvaluatedAt(OffsetDateTime.parse("2026-10-15T10:00:00Z")),
            com.azure.core.util.Context.NONE);
    }
}
