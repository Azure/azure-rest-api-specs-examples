import com.azure.resourcemanager.costmanagement.models.ExportDataset;
import com.azure.resourcemanager.costmanagement.models.ExportDatasetConfiguration;
import com.azure.resourcemanager.costmanagement.models.ExportDefinition;
import com.azure.resourcemanager.costmanagement.models.ExportDeliveryDestination;
import com.azure.resourcemanager.costmanagement.models.ExportDeliveryInfo;
import com.azure.resourcemanager.costmanagement.models.ExportRecurrencePeriod;
import com.azure.resourcemanager.costmanagement.models.ExportSchedule;
import com.azure.resourcemanager.costmanagement.models.ExportType;
import com.azure.resourcemanager.costmanagement.models.FormatType;
import com.azure.resourcemanager.costmanagement.models.GranularityType;
import com.azure.resourcemanager.costmanagement.models.RecurrenceType;
import com.azure.resourcemanager.costmanagement.models.StatusType;
import com.azure.resourcemanager.costmanagement.models.TimeframeType;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for Exports CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportCreateOrUpdateByResourceGroup.json
     */
    /**
     * Sample code: ExportCreateOrUpdateByResourceGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportCreateOrUpdateByResourceGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .define("TestExport")
            .withExistingScope("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG")
            .withSchedule(
                new ExportSchedule()
                    .withStatus(StatusType.ACTIVE)
                    .withRecurrence(RecurrenceType.WEEKLY)
                    .withRecurrencePeriod(
                        new ExportRecurrencePeriod()
                            .withFrom(OffsetDateTime.parse("2020-06-01T00:00:00Z"))
                            .withTo(OffsetDateTime.parse("2020-10-31T00:00:00Z"))))
            .withFormat(FormatType.CSV)
            .withDeliveryInfo(
                new ExportDeliveryInfo()
                    .withDestination(
                        new ExportDeliveryDestination()
                            .withResourceId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182")
                            .withContainer("exports")
                            .withRootFolderPath("ad-hoc")))
            .withDefinition(
                new ExportDefinition()
                    .withType(ExportType.ACTUAL_COST)
                    .withTimeframe(TimeframeType.MONTH_TO_DATE)
                    .withDataSet(
                        new ExportDataset()
                            .withGranularity(GranularityType.DAILY)
                            .withConfiguration(
                                new ExportDatasetConfiguration()
                                    .withColumns(
                                        Arrays
                                            .asList("Date", "MeterId", "ResourceId", "ResourceLocation", "Quantity")))))
            .create();
    }
}
