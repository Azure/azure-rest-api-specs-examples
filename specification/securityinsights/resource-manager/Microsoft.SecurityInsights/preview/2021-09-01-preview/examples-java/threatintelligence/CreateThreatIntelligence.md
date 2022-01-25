Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.ThreatIntelligenceIndicatorModelForRequestBody;
import com.azure.resourcemanager.securityinsights.models.ThreatIntelligenceResourceKindEnum;
import java.util.Arrays;

/** Samples for ThreatIntelligenceIndicator CreateIndicator. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/threatintelligence/CreateThreatIntelligence.json
     */
    /**
     * Sample code: Create a new Threat Intelligence.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createANewThreatIntelligence(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .threatIntelligenceIndicators()
            .createIndicatorWithResponse(
                "myRg",
                "myWorkspace",
                new ThreatIntelligenceIndicatorModelForRequestBody()
                    .withKind(ThreatIntelligenceResourceKindEnum.INDICATOR)
                    .withThreatIntelligenceTags(Arrays.asList("new schema"))
                    .withSource("Azure Sentinel")
                    .withDisplayName("new schema")
                    .withDescription("debugging indicators")
                    .withPattern("[url:value = 'https://www.contoso.com']")
                    .withPatternType("url")
                    .withKillChainPhases(Arrays.asList())
                    .withCreatedByRef("contoso@contoso.com")
                    .withExternalReferences(Arrays.asList())
                    .withGranularMarkings(Arrays.asList())
                    .withLabels(Arrays.asList())
                    .withRevoked(false)
                    .withConfidence(78)
                    .withThreatTypes(Arrays.asList("compromised"))
                    .withValidFrom("2021-09-15T17:44:00.114052Z")
                    .withValidUntil("")
                    .withModified(""),
                Context.NONE);
    }
}
```
