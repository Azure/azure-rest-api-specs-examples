
import com.azure.resourcemanager.consumption.models.BudgetComparisonExpression;
import com.azure.resourcemanager.consumption.models.BudgetFilter;
import com.azure.resourcemanager.consumption.models.BudgetFilterProperties;
import com.azure.resourcemanager.consumption.models.BudgetOperatorType;
import com.azure.resourcemanager.consumption.models.BudgetTimePeriod;
import com.azure.resourcemanager.consumption.models.CategoryType;
import com.azure.resourcemanager.consumption.models.CultureCode;
import com.azure.resourcemanager.consumption.models.Notification;
import com.azure.resourcemanager.consumption.models.OperatorType;
import com.azure.resourcemanager.consumption.models.ThresholdType;
import com.azure.resourcemanager.consumption.models.TimeGrainType;
import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Budgets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/CreateOrUpdateBudget.
     * json
     */
    /**
     * Sample code: CreateOrUpdateBudget.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void createOrUpdateBudget(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.budgets().define("TestBudget").withExistingScope("subscriptions/00000000-0000-0000-0000-000000000000")
            .withEtag("\"1d34d016a593709\"").withCategory(CategoryType.COST).withAmount(new BigDecimal("100.65"))
            .withTimeGrain(TimeGrainType.MONTHLY)
            .withTimePeriod(new BudgetTimePeriod().withStartDate(OffsetDateTime.parse("2017-10-01T00:00:00Z"))
                .withEndDate(OffsetDateTime.parse("2018-10-31T00:00:00Z")))
            .withFilter(new BudgetFilter().withAnd(Arrays.asList(
                new BudgetFilterProperties().withDimensions(new BudgetComparisonExpression().withName("ResourceId")
                    .withOperator(BudgetOperatorType.IN)
                    .withValues(Arrays.asList(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2",
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1"))),
                new BudgetFilterProperties().withTags(new BudgetComparisonExpression().withName("category")
                    .withOperator(BudgetOperatorType.IN).withValues(Arrays.asList("Dev", "Prod"))),
                new BudgetFilterProperties().withTags(new BudgetComparisonExpression().withName("department")
                    .withOperator(BudgetOperatorType.IN).withValues(Arrays.asList("engineering", "sales"))))))
            .withNotifications(mapOf("Actual_GreaterThan_80_Percent", new Notification().withEnabled(true)
                .withOperator(OperatorType.GREATER_THAN).withThreshold(new BigDecimal("80"))
                .withContactEmails(Arrays.asList("johndoe@contoso.com", "janesmith@contoso.com"))
                .withContactRoles(Arrays.asList("Contributor", "Reader"))
                .withContactGroups(Arrays.asList(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup"))
                .withThresholdType(ThresholdType.ACTUAL).withLocale(CultureCode.EN_US)))
            .create();
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
