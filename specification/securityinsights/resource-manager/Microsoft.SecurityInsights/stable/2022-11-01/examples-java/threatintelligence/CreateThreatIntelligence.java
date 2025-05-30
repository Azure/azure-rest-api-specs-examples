
import com.azure.resourcemanager.securityinsights.models.ThreatIntelligenceIndicatorModel;
import java.util.Arrays;

/**
 * Samples for ThreatIntelligenceIndicator CreateIndicator.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/
     * threatintelligence/CreateThreatIntelligence.json
     */
    /**
     * Sample code: Create a new Threat Intelligence.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        createANewThreatIntelligence(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.threatIntelligenceIndicators().createIndicatorWithResponse("myRg", "myWorkspace",
            new ThreatIntelligenceIndicatorModel().withThreatIntelligenceTags(Arrays.asList("new schema"))
                .withSource("Azure Sentinel").withDisplayName("new schema").withDescription("debugging indicators")
                .withPattern("[url:value = 'https://www.contoso.com']").withPatternType("url")
                .withKillChainPhases(Arrays.asList()).withCreatedByRef("contoso@contoso.com")
                .withExternalReferences(Arrays.asList()).withGranularMarkings(Arrays.asList())
                .withLabels(Arrays.asList()).withRevoked(false).withConfidence(78)
                .withThreatTypes(Arrays.asList("compromised")).withValidFrom("2020-04-15T17:44:00.114052Z")
                .withValidUntil("").withModified(""),
            com.azure.core.util.Context.NONE);
    }
}
