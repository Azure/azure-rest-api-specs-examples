
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
     * x-ms-original-file: 2025-04-01/queryCarbonEmissionsMonthlySummaryReportWithOtherOptionalFilter.json
     */
    /**
     * Sample code: QueryCarbonEmission Monthly Summary Report with optional filter - locationList, resourceTypeList,
     * resourceGroupUrlList.
     * 
     * @param manager Entry point to CarbonOptimizationManager.
     */
    public static void
        queryCarbonEmissionMonthlySummaryReportWithOptionalFilterLocationListResourceTypeListResourceGroupUrlList(
            com.azure.resourcemanager.carbonoptimization.CarbonOptimizationManager manager) {
        manager.carbonServices().queryCarbonEmissionReportsWithResponse(
            new MonthlySummaryReportQueryFilter()
                .withDateRange(
                    new DateRange().withStart(LocalDate.parse("2024-03-01")).withEnd(LocalDate.parse("2024-05-01")))
                .withSubscriptionList(Arrays.asList("00000000-0000-0000-0000-000000000000"))
                .withResourceGroupUrlList(
                    Arrays.asList("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-name"))
                .withResourceTypeList(
                    Arrays.asList("microsoft.storage/storageaccounts", "microsoft.databricks/workspaces"))
                .withLocationList(Arrays.asList("east us", "west us"))
                .withCarbonScopeList(Arrays.asList(EmissionScopeEnum.SCOPE1, EmissionScopeEnum.SCOPE3)),
            com.azure.core.util.Context.NONE);
    }
}
