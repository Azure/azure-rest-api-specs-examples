Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.2/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.ThreatIntelligenceIndicatorModel;
import java.util.Arrays;

/** Samples for ThreatIntelligenceIndicator ReplaceTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/threatintelligence/ReplaceTagsThreatIntelligence.json
     */
    /**
     * Sample code: Replace tags to a Threat Intelligence.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void replaceTagsToAThreatIntelligence(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .threatIntelligenceIndicators()
            .replaceTagsWithResponse(
                "myRg",
                "myWorkspace",
                "d9cd6f0b-96b9-3984-17cd-a779d1e15a93",
                new ThreatIntelligenceIndicatorModel()
                    .withEtag("\"0000262c-0000-0800-0000-5e9767060000\"")
                    .withThreatIntelligenceTags(Arrays.asList("patching tags")),
                Context.NONE);
    }
}
```
