import com.azure.resourcemanager.frontdoor.models.TimeseriesAggregationInterval;
import com.azure.resourcemanager.frontdoor.models.TimeseriesType;
import java.time.OffsetDateTime;

/** Samples for Reports GetTimeseries. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetTimeseries.json
     */
    /**
     * Sample code: Gets a Timeseries for a given Experiment.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getsATimeseriesForAGivenExperiment(
        com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager
            .reports()
            .getTimeseriesWithResponse(
                "MyResourceGroup",
                "MyProfile",
                "MyExperiment",
                OffsetDateTime.parse("2019-07-21T17:32:28Z"),
                OffsetDateTime.parse("2019-09-21T17:32:28Z"),
                TimeseriesAggregationInterval.HOURLY,
                TimeseriesType.MEASUREMENT_COUNTS,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
