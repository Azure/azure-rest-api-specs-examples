import com.azure.resourcemanager.costmanagement.models.ExportType;
import com.azure.resourcemanager.costmanagement.models.GranularityType;
import com.azure.resourcemanager.costmanagement.models.QueryComparisonExpression;
import com.azure.resourcemanager.costmanagement.models.QueryDataset;
import com.azure.resourcemanager.costmanagement.models.QueryDefinition;
import com.azure.resourcemanager.costmanagement.models.QueryFilter;
import com.azure.resourcemanager.costmanagement.models.QueryOperatorType;
import com.azure.resourcemanager.costmanagement.models.TimeframeType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Query Usage. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/MCAInvoiceSectionQuery.json
     */
    /**
     * Sample code: InvoiceSectionQuery-Modern.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void invoiceSectionQueryModern(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .queries()
            .usageWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876",
                new QueryDefinition()
                    .withType(ExportType.USAGE)
                    .withTimeframe(TimeframeType.MONTH_TO_DATE)
                    .withDataset(
                        new QueryDataset()
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
                                                            .withValues(Arrays.asList("API"))))))),
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
