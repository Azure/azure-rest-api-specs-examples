
import com.azure.resourcemanager.costmanagement.models.ForecastAggregation;
import com.azure.resourcemanager.costmanagement.models.ForecastComparisonExpression;
import com.azure.resourcemanager.costmanagement.models.ForecastDataset;
import com.azure.resourcemanager.costmanagement.models.ForecastDefinition;
import com.azure.resourcemanager.costmanagement.models.ForecastFilter;
import com.azure.resourcemanager.costmanagement.models.ForecastOperatorType;
import com.azure.resourcemanager.costmanagement.models.ForecastTimePeriod;
import com.azure.resourcemanager.costmanagement.models.ForecastTimeframe;
import com.azure.resourcemanager.costmanagement.models.ForecastType;
import com.azure.resourcemanager.costmanagement.models.FunctionName;
import com.azure.resourcemanager.costmanagement.models.FunctionType;
import com.azure.resourcemanager.costmanagement.models.GranularityType;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Forecast Usage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * BillingProfileForecast.json
     */
    /**
     * Sample code: BillingProfileForecast.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void billingProfileForecast(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.forecasts().usageWithResponse(
            "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579",
            new ForecastDefinition().withType(ForecastType.USAGE).withTimeframe(ForecastTimeframe.CUSTOM)
                .withTimePeriod(new ForecastTimePeriod().withFrom(OffsetDateTime.parse("2022-08-01T00:00:00+00:00"))
                    .withTo(OffsetDateTime.parse("2022-08-31T23:59:59+00:00")))
                .withDataset(new ForecastDataset().withGranularity(GranularityType.DAILY)
                    .withAggregation(mapOf("totalCost",
                        new ForecastAggregation().withName(FunctionName.COST).withFunction(FunctionType.SUM)))
                    .withFilter(new ForecastFilter().withAnd(Arrays.asList(
                        new ForecastFilter().withOr(Arrays.asList(
                            new ForecastFilter().withDimensions(new ForecastComparisonExpression()
                                .withName("ResourceLocation").withOperator(ForecastOperatorType.IN)
                                .withValues(Arrays.asList("East US", "West Europe"))),
                            new ForecastFilter().withTags(new ForecastComparisonExpression().withName("Environment")
                                .withOperator(ForecastOperatorType.IN).withValues(Arrays.asList("UAT", "Prod"))))),
                        new ForecastFilter().withDimensions(new ForecastComparisonExpression().withName("ResourceGroup")
                            .withOperator(ForecastOperatorType.IN).withValues(Arrays.asList("API")))))))
                .withIncludeActualCost(false).withIncludeFreshPartialCost(false),
            null, com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
