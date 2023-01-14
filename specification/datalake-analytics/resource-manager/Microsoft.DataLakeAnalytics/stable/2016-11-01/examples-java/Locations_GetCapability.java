/** Samples for Locations GetCapability. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/Locations_GetCapability.json
     */
    /**
     * Sample code: Gets subscription-level properties and limits for Data Lake Analytics specified by resource
     * location.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void getsSubscriptionLevelPropertiesAndLimitsForDataLakeAnalyticsSpecifiedByResourceLocation(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.locations().getCapabilityWithResponse("EastUS2", com.azure.core.util.Context.NONE);
    }
}
