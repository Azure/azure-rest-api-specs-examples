import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.models.TopQueryStatisticsInput;
import java.time.OffsetDateTime;

/** Samples for TopQueryStatistics ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/TopQueryStatisticsListByServer.json
     */
    /**
     * Sample code: TopQueryStatisticsListByServer.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void topQueryStatisticsListByServer(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .topQueryStatistics()
            .listByServer(
                "testResourceGroupName",
                "testServerName",
                new TopQueryStatisticsInput()
                    .withNumberOfTopQueries(5)
                    .withAggregationFunction("avg")
                    .withObservedMetric("duration")
                    .withObservationStartTime(OffsetDateTime.parse("2019-05-01T20:00:00.000Z"))
                    .withObservationEndTime(OffsetDateTime.parse("2019-05-07T20:00:00.000Z"))
                    .withAggregationWindow("PT15M"),
                Context.NONE);
    }
}
