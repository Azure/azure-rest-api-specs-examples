```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.ThreatIntelligenceFilteringCriteria;
import com.azure.resourcemanager.securityinsights.models.ThreatIntelligenceSortingCriteria;
import com.azure.resourcemanager.securityinsights.models.ThreatIntelligenceSortingCriteriaEnum;
import java.util.Arrays;

/** Samples for ThreatIntelligenceIndicator QueryIndicators. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/threatintelligence/QueryThreatIntelligence.json
     */
    /**
     * Sample code: Query threat intelligence indicators as per filtering criteria.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void queryThreatIntelligenceIndicatorsAsPerFilteringCriteria(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .threatIntelligenceIndicators()
            .queryIndicators(
                "myRg",
                "myWorkspace",
                new ThreatIntelligenceFilteringCriteria()
                    .withPageSize(100)
                    .withMinConfidence(25)
                    .withMaxConfidence(80)
                    .withMinValidUntil("2021-04-05T17:44:00.114052Z")
                    .withMaxValidUntil("2021-04-25T17:44:00.114052Z")
                    .withSortBy(
                        Arrays
                            .asList(
                                new ThreatIntelligenceSortingCriteria()
                                    .withItemKey("lastUpdatedTimeUtc")
                                    .withSortOrder(ThreatIntelligenceSortingCriteriaEnum.DESCENDING)))
                    .withSources(Arrays.asList("Azure Sentinel")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
