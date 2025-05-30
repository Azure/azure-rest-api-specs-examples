
import com.azure.resourcemanager.costmanagement.models.ExportType;
import com.azure.resourcemanager.costmanagement.models.ExternalCloudProviderType;
import com.azure.resourcemanager.costmanagement.models.GranularityType;
import com.azure.resourcemanager.costmanagement.models.QueryComparisonExpression;
import com.azure.resourcemanager.costmanagement.models.QueryDataset;
import com.azure.resourcemanager.costmanagement.models.QueryDefinition;
import com.azure.resourcemanager.costmanagement.models.QueryFilter;
import com.azure.resourcemanager.costmanagement.models.QueryOperatorType;
import com.azure.resourcemanager.costmanagement.models.TimeframeType;
import java.util.Arrays;

/**
 * Samples for Query UsageByExternalCloudProviderType.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * ExternalBillingAccountsQuery.json
     */
    /**
     * Sample code: ExternalBillingAccountQueryList.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void
        externalBillingAccountQueryList(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.queries().usageByExternalCloudProviderTypeWithResponse(
            ExternalCloudProviderType.EXTERNAL_BILLING_ACCOUNTS, "100",
            new QueryDefinition().withType(ExportType.USAGE).withTimeframe(TimeframeType.MONTH_TO_DATE)
                .withDataset(new QueryDataset().withGranularity(GranularityType.DAILY)
                    .withFilter(new QueryFilter().withAnd(Arrays.asList(new QueryFilter().withOr(Arrays.asList(
                        new QueryFilter().withDimensions(new QueryComparisonExpression().withName("ResourceLocation")
                            .withOperator(QueryOperatorType.IN).withValues(Arrays.asList("East US", "West Europe"))),
                        new QueryFilter().withTags(new QueryComparisonExpression().withName("Environment")
                            .withOperator(QueryOperatorType.IN).withValues(Arrays.asList("UAT", "Prod"))))),
                        new QueryFilter().withDimensions(new QueryComparisonExpression().withName("ResourceGroup")
                            .withOperator(QueryOperatorType.IN).withValues(Arrays.asList("API"))))))),
            com.azure.core.util.Context.NONE);
    }
}
