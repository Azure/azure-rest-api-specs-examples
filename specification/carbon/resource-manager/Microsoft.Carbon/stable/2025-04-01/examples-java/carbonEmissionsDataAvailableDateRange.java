
/**
 * Samples for CarbonService QueryCarbonEmissionDataAvailableDateRange.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/carbonEmissionsDataAvailableDateRange.json
     */
    /**
     * Sample code: CarbonService_QueryCarbonEmissionDataAvailableDateRange.
     * 
     * @param manager Entry point to CarbonOptimizationManager.
     */
    public static void carbonServiceQueryCarbonEmissionDataAvailableDateRange(
        com.azure.resourcemanager.carbonoptimization.CarbonOptimizationManager manager) {
        manager.carbonServices()
            .queryCarbonEmissionDataAvailableDateRangeWithResponse(com.azure.core.util.Context.NONE);
    }
}
