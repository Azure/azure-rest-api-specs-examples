
import com.azure.resourcemanager.loganalytics.models.Column;
import com.azure.resourcemanager.loganalytics.models.ColumnTypeEnum;
import com.azure.resourcemanager.loganalytics.models.Schema;
import com.azure.resourcemanager.loganalytics.models.Table;
import java.util.Arrays;

/**
 * Samples for Tables Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2022-10-01/examples/
     * TablesUpsert.json
     */
    /**
     * Sample code: TablesUpsert.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesUpsert(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        Table resource = manager.tables()
            .getWithResponse("oiautorest6685", "oiautorest6685", "AzureNetworkFlow", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withRetentionInDays(45).withTotalRetentionInDays(70)
            .withSchema(new Schema().withName("AzureNetworkFlow")
                .withColumns(Arrays.asList(new Column().withName("MyNewColumn").withType(ColumnTypeEnum.GUID))))
            .apply();
    }
}
