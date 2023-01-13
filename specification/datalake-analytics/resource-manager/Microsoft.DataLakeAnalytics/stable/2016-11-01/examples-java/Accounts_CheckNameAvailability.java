import com.azure.resourcemanager.datalakeanalytics.models.CheckNameAvailabilityParameters;

/** Samples for Accounts CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/Accounts_CheckNameAvailability.json
     */
    /**
     * Sample code: Checks whether the specified account name is available or taken.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void checksWhetherTheSpecifiedAccountNameIsAvailableOrTaken(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .accounts()
            .checkNameAvailabilityWithResponse(
                "EastUS2",
                new CheckNameAvailabilityParameters().withName("contosoadla"),
                com.azure.core.util.Context.NONE);
    }
}
