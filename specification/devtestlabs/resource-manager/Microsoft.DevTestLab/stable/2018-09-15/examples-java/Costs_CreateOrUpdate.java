
import com.azure.resourcemanager.devtestlabs.models.CostThresholdProperties;
import com.azure.resourcemanager.devtestlabs.models.CostThresholdStatus;
import com.azure.resourcemanager.devtestlabs.models.PercentageCostThresholdProperties;
import com.azure.resourcemanager.devtestlabs.models.ReportingCycleType;
import com.azure.resourcemanager.devtestlabs.models.TargetCostProperties;
import com.azure.resourcemanager.devtestlabs.models.TargetCostStatus;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for Costs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_CreateOrUpdate.
     * json
     */
    /**
     * Sample code: Costs_CreateOrUpdate.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void costsCreateOrUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .costs().define(
                "targetCost")
            .withRegion((String) null).withExistingLab("resourceGroupName",
                "{labName}")
            .withTargetCost(new TargetCostProperties().withStatus(TargetCostStatus.ENABLED).withTarget(100)
                .withCostThresholds(Arrays.asList(
                    new CostThresholdProperties().withThresholdId("00000000-0000-0000-0000-000000000001")
                        .withPercentageThreshold(new PercentageCostThresholdProperties().withThresholdValue(25.0D))
                        .withDisplayOnChart(CostThresholdStatus.DISABLED)
                        .withSendNotificationWhenExceeded(CostThresholdStatus.DISABLED),
                    new CostThresholdProperties().withThresholdId("00000000-0000-0000-0000-000000000002")
                        .withPercentageThreshold(new PercentageCostThresholdProperties().withThresholdValue(50.0D))
                        .withDisplayOnChart(CostThresholdStatus.ENABLED)
                        .withSendNotificationWhenExceeded(CostThresholdStatus.ENABLED),
                    new CostThresholdProperties().withThresholdId("00000000-0000-0000-0000-000000000003")
                        .withPercentageThreshold(new PercentageCostThresholdProperties().withThresholdValue(75.0D))
                        .withDisplayOnChart(CostThresholdStatus.DISABLED)
                        .withSendNotificationWhenExceeded(CostThresholdStatus.DISABLED),
                    new CostThresholdProperties().withThresholdId("00000000-0000-0000-0000-000000000004")
                        .withPercentageThreshold(new PercentageCostThresholdProperties().withThresholdValue(100.0D))
                        .withDisplayOnChart(CostThresholdStatus.DISABLED)
                        .withSendNotificationWhenExceeded(CostThresholdStatus.DISABLED),
                    new CostThresholdProperties().withThresholdId("00000000-0000-0000-0000-000000000005")
                        .withPercentageThreshold(new PercentageCostThresholdProperties().withThresholdValue(125.0D))
                        .withDisplayOnChart(CostThresholdStatus.DISABLED)
                        .withSendNotificationWhenExceeded(CostThresholdStatus.DISABLED)))
                .withCycleStartDateTime(OffsetDateTime.parse("2020-12-01T00:00:00.000Z"))
                .withCycleEndDateTime(OffsetDateTime.parse("2020-12-31T00:00:00.000Z"))
                .withCycleType(ReportingCycleType.CALENDAR_MONTH))
            .withCurrencyCode("USD").withStartDateTime(OffsetDateTime.parse("2020-12-01T00:00:00Z"))
            .withEndDateTime(OffsetDateTime.parse("2020-12-31T23:59:59Z")).create();
    }
}
