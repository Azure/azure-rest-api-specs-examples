```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.LogRanking;
import com.azure.resourcemanager.cdn.models.LogRankingMetric;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for LogAnalytics GetLogAnalyticsRankings. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsRankings.json
     */
    /**
     * Sample code: LogAnalytics_GetLogAnalyticsRankings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void logAnalyticsGetLogAnalyticsRankings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getLogAnalytics()
            .getLogAnalyticsRankingsWithResponse(
                "RG",
                "profile1",
                Arrays.asList(LogRanking.URL),
                Arrays.asList(LogRankingMetric.CLIENT_REQUEST_COUNT),
                5,
                OffsetDateTime.parse("2020-11-04T06:49:27.554Z"),
                OffsetDateTime.parse("2020-11-04T09:49:27.554Z"),
                null,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
