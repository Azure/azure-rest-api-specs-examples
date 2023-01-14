import com.azure.resourcemanager.costmanagement.models.ExportType;
import com.azure.resourcemanager.costmanagement.models.FunctionType;
import com.azure.resourcemanager.costmanagement.models.GranularityType;
import com.azure.resourcemanager.costmanagement.models.QueryAggregation;
import com.azure.resourcemanager.costmanagement.models.QueryColumnType;
import com.azure.resourcemanager.costmanagement.models.QueryDataset;
import com.azure.resourcemanager.costmanagement.models.QueryDefinition;
import com.azure.resourcemanager.costmanagement.models.QueryGrouping;
import com.azure.resourcemanager.costmanagement.models.TimeframeType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Query Usage. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ResourceGroupQueryGrouping.json
     */
    /**
     * Sample code: ResourceGroupQueryGrouping-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void resourceGroupQueryGroupingLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .queries()
            .usageWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer",
                new QueryDefinition()
                    .withType(ExportType.USAGE)
                    .withTimeframe(TimeframeType.THE_LAST_MONTH)
                    .withDataset(
                        new QueryDataset()
                            .withGranularity(GranularityType.DAILY)
                            .withAggregation(
                                mapOf(
                                    "totalCost",
                                    new QueryAggregation().withName("PreTaxCost").withFunction(FunctionType.SUM)))
                            .withGrouping(
                                Arrays
                                    .asList(
                                        new QueryGrouping()
                                            .withType(QueryColumnType.DIMENSION)
                                            .withName("ResourceType")))),
                com.azure.core.util.Context.NONE);
    }

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
