
import com.azure.resourcemanager.carbonoptimization.models.DateRange;
import com.azure.resourcemanager.carbonoptimization.models.EmissionScopeEnum;
import com.azure.resourcemanager.carbonoptimization.models.MonthlySummaryReportQueryFilter;
import java.time.LocalDate;
import java.util.Arrays;

/**
 * Samples for CarbonService QueryCarbonEmissionReports.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/queryCarbonEmissionsMonthlySummaryReport.json
     */
    /**
     * Sample code: QueryCarbonEmission Overall Monthly Summary Report.
     * 
     * @param manager Entry point to CarbonOptimizationManager.
     */
    public static void queryCarbonEmissionOverallMonthlySummaryReport(
        com.azure.resourcemanager.carbonoptimization.CarbonOptimizationManager manager) {
        manager.carbonServices().queryCarbonEmissionReportsWithResponse(
            new MonthlySummaryReportQueryFilter()
                .withDateRange(
                    new DateRange().withStart(LocalDate.parse("2024-03-01")).withEnd(LocalDate.parse("2024-05-01")))
                .withSubscriptionList(Arrays.asList("00000000-0000-0000-0000-000000000000"))
                .withCarbonScopeList(Arrays.asList(EmissionScopeEnum.SCOPE1, EmissionScopeEnum.SCOPE3)),
            com.azure.core.util.Context.NONE);
    }
}
