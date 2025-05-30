
import com.azure.resourcemanager.costmanagement.fluent.models.ViewInner;
import com.azure.resourcemanager.costmanagement.models.AccumulatedType;
import com.azure.resourcemanager.costmanagement.models.ChartType;
import com.azure.resourcemanager.costmanagement.models.FunctionType;
import com.azure.resourcemanager.costmanagement.models.KpiProperties;
import com.azure.resourcemanager.costmanagement.models.KpiType;
import com.azure.resourcemanager.costmanagement.models.MetricType;
import com.azure.resourcemanager.costmanagement.models.PivotProperties;
import com.azure.resourcemanager.costmanagement.models.PivotType;
import com.azure.resourcemanager.costmanagement.models.ReportConfigAggregation;
import com.azure.resourcemanager.costmanagement.models.ReportConfigDataset;
import com.azure.resourcemanager.costmanagement.models.ReportConfigSorting;
import com.azure.resourcemanager.costmanagement.models.ReportConfigSortingType;
import com.azure.resourcemanager.costmanagement.models.ReportGranularityType;
import com.azure.resourcemanager.costmanagement.models.ReportTimeframeType;
import com.azure.resourcemanager.costmanagement.models.ReportType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Views CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * PrivateViewCreateOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdatePrivateView.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void
        createOrUpdatePrivateView(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.views().createOrUpdateWithResponse("swaggerExample", new ViewInner().withEtag("\"1d4ff9fe66f1d10\"")
            .withDisplayName("swagger Example").withChart(ChartType.TABLE).withAccumulated(AccumulatedType.TRUE)
            .withMetric(MetricType.ACTUAL_COST)
            .withKpis(Arrays.asList(new KpiProperties().withType(KpiType.FORECAST).withEnabled(true),
                new KpiProperties().withType(KpiType.BUDGET).withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Consumption/budgets/swaggerDemo")
                    .withEnabled(true)))
            .withPivots(Arrays.asList(new PivotProperties().withType(PivotType.DIMENSION).withName("ServiceName"),
                new PivotProperties().withType(PivotType.DIMENSION).withName("MeterCategory"),
                new PivotProperties().withType(PivotType.TAG_KEY).withName("swaggerTagKey")))
            .withTypePropertiesType(ReportType.USAGE).withTimeframe(ReportTimeframeType.MONTH_TO_DATE)
            .withDataSet(new ReportConfigDataset().withGranularity(ReportGranularityType.DAILY)
                .withAggregation(mapOf("totalCost",
                    new ReportConfigAggregation().withName("PreTaxCost").withFunction(FunctionType.SUM)))
                .withGrouping(Arrays.asList())
                .withSorting(Arrays.asList(
                    new ReportConfigSorting().withDirection(ReportConfigSortingType.ASCENDING).withName("UsageDate")))),
            com.azure.core.util.Context.NONE);
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
