import com.azure.resourcemanager.frontdoor.models.LatencyScorecardAggregationInterval;

/** Samples for Reports GetLatencyScorecards. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetLatencyScorecard.json
     */
    /**
     * Sample code: Gets a Latency Scorecard for a given Experiment.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getsALatencyScorecardForAGivenExperiment(
        com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager
            .reports()
            .getLatencyScorecardsWithResponse(
                "MyResourceGroup",
                "MyProfile",
                "MyExperiment",
                LatencyScorecardAggregationInterval.DAILY,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
