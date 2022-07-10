import com.azure.resourcemanager.loganalytics.models.Column;
import com.azure.resourcemanager.loganalytics.models.ColumnTypeEnum;
import com.azure.resourcemanager.loganalytics.models.Schema;
import java.util.Arrays;

/** Samples for Tables CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/TablesUpsert.json
     */
    /**
     * Sample code: TablesUpsert.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesUpsert(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .tables()
            .define("AzureNetworkFlow")
            .withExistingWorkspace("oiautorest6685", "oiautorest6685")
            .withRetentionInDays(45)
            .withTotalRetentionInDays(70)
            .withSchema(
                new Schema()
                    .withName("AzureNetworkFlow")
                    .withColumns(Arrays.asList(new Column().withName("MyNewColumn").withType(ColumnTypeEnum.GUID))))
            .create();
    }
}
