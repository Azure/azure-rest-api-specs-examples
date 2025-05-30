
import com.azure.resourcemanager.carbonoptimization.models.CategoryTypeEnum;
import com.azure.resourcemanager.carbonoptimization.models.DateRange;
import com.azure.resourcemanager.carbonoptimization.models.EmissionScopeEnum;
import com.azure.resourcemanager.carbonoptimization.models.TopItemsMonthlySummaryReportQueryFilter;
import java.time.LocalDate;
import java.util.Arrays;

/**
 * Samples for CarbonService QueryCarbonEmissionReports.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/queryCarbonEmissionsTopNResourceTypeItemsMonthlyReport.json
     */
    /**
     * Sample code: QueryCarbonEmission Top N ResourceType monthly Report.
     * 
     * @param manager Entry point to CarbonOptimizationManager.
     */
    public static void queryCarbonEmissionTopNResourceTypeMonthlyReport(
        com.azure.resourcemanager.carbonoptimization.CarbonOptimizationManager manager) {
        manager.carbonServices()
            .queryCarbonEmissionReportsWithResponse(
                new TopItemsMonthlySummaryReportQueryFilter()
                    .withDateRange(
                        new DateRange().withStart(LocalDate.parse("2024-03-01")).withEnd(LocalDate.parse("2024-05-01")))
                    .withSubscriptionList(Arrays.asList("00000000-0000-0000-0000-000000000000",
                        "00000000-0000-0000-0000-000000000001,", "00000000-0000-0000-0000-000000000002",
                        "00000000-0000-0000-0000-000000000003", "00000000-0000-0000-0000-000000000004",
                        "00000000-0000-0000-0000-000000000005", "00000000-0000-0000-0000-000000000006",
                        "00000000-0000-0000-0000-000000000007", "00000000-0000-0000-0000-000000000008"))
                    .withCarbonScopeList(Arrays.asList(EmissionScopeEnum.SCOPE1, EmissionScopeEnum.SCOPE3))
                    .withCategoryType(CategoryTypeEnum.RESOURCE_TYPE).withTopItems(2),
                com.azure.core.util.Context.NONE);
    }
}
