
import com.azure.resourcemanager.commerce.models.AggregationGranularity;
import java.time.OffsetDateTime;

/**
 * Samples for UsageAggregates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2015-06-01-preview/GetUsageAggregatesList.json
     */
    /**
     * Sample code: GetUsageAggregatesList.
     * 
     * @param manager Entry point to CommerceManager.
     */
    public static void getUsageAggregatesList(com.azure.resourcemanager.commerce.CommerceManager manager) {
        manager.usageAggregates().list(OffsetDateTime.parse("2014-05-01T00:00:00 00:00"),
            OffsetDateTime.parse("2015-06-01T00:00:00 00:00"), true, AggregationGranularity.DAILY, null,
            com.azure.core.util.Context.NONE);
    }
}
