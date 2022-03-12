Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.WafAction;
import com.azure.resourcemanager.cdn.models.WafGranularity;
import com.azure.resourcemanager.cdn.models.WafMetric;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for LogAnalytics GetWafLogAnalyticsMetrics. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetWafLogAnalyticsMetrics.json
     */
    /**
     * Sample code: LogAnalytics_GetWafLogAnalyticsMetrics.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void logAnalyticsGetWafLogAnalyticsMetrics(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getLogAnalytics()
            .getWafLogAnalyticsMetricsWithResponse(
                "RG",
                "profile1",
                Arrays.asList(WafMetric.CLIENT_REQUEST_COUNT),
                OffsetDateTime.parse("2020-11-04T06:49:27.554Z"),
                OffsetDateTime.parse("2020-11-04T09:49:27.554Z"),
                WafGranularity.PT5M,
                Arrays.asList(WafAction.BLOCK, WafAction.LOG),
                null,
                null,
                Context.NONE);
    }
}
```
