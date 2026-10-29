
import com.azure.resourcemanager.cdn.models.WafMetric;
import com.azure.resourcemanager.cdn.models.WafRankingType;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for LogAnalytics GetWafLogAnalyticsRankings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/LogAnalytics_GetWafLogAnalyticsRankings.json
     */
    /**
     * Sample code: LogAnalytics_GetWafLogAnalyticsRankings.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void logAnalyticsGetWafLogAnalyticsRankings(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getLogAnalytics().getWafLogAnalyticsRankingsWithResponse("RG", "profile1",
            Arrays.asList(WafMetric.CLIENT_REQUEST_COUNT), OffsetDateTime.parse("2020-11-04T06:49:27.554Z"),
            OffsetDateTime.parse("2020-11-04T09:49:27.554Z"), 5, Arrays.asList(WafRankingType.RULE_ID), null, null,
            com.azure.core.util.Context.NONE);
    }
}
