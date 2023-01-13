import com.azure.resourcemanager.costmanagement.models.ForecastDataset;
import com.azure.resourcemanager.costmanagement.models.ForecastDefinition;
import com.azure.resourcemanager.costmanagement.models.ForecastTimeframeType;
import com.azure.resourcemanager.costmanagement.models.ForecastType;
import com.azure.resourcemanager.costmanagement.models.GranularityType;
import com.azure.resourcemanager.costmanagement.models.QueryComparisonExpression;
import com.azure.resourcemanager.costmanagement.models.QueryFilter;
import com.azure.resourcemanager.costmanagement.models.QueryOperatorType;
import java.util.Arrays;

/** Samples for Forecast Usage. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/SubscriptionForecast.json
     */
    /**
     * Sample code: SubscriptionForecast.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void subscriptionForecast(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .forecasts()
            .usageWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                new ForecastDefinition()
                    .withType(ForecastType.USAGE)
                    .withTimeframe(ForecastTimeframeType.MONTH_TO_DATE)
                    .withDataset(
                        new ForecastDataset()
                            .withGranularity(GranularityType.DAILY)
                            .withFilter(
                                new QueryFilter()
                                    .withAnd(
                                        Arrays
                                            .asList(
                                                new QueryFilter()
                                                    .withOr(
                                                        Arrays
                                                            .asList(
                                                                new QueryFilter()
                                                                    .withDimension(
                                                                        new QueryComparisonExpression()
                                                                            .withName("ResourceLocation")
                                                                            .withOperator(QueryOperatorType.IN)
                                                                            .withValues(
                                                                                Arrays
                                                                                    .asList("East US", "West Europe"))),
                                                                new QueryFilter()
                                                                    .withTag(
                                                                        new QueryComparisonExpression()
                                                                            .withName("Environment")
                                                                            .withOperator(QueryOperatorType.IN)
                                                                            .withValues(
                                                                                Arrays.asList("UAT", "Prod"))))),
                                                new QueryFilter()
                                                    .withDimension(
                                                        new QueryComparisonExpression()
                                                            .withName("ResourceGroup")
                                                            .withOperator(QueryOperatorType.IN)
                                                            .withValues(Arrays.asList("API")))))))
                    .withIncludeActualCost(false)
                    .withIncludeFreshPartialCost(false),
                null,
                com.azure.core.util.Context.NONE);
    }
}
