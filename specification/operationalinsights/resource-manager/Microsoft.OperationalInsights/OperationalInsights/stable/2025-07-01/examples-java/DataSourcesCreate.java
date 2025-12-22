
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.loganalytics.models.DataSourceKind;
import java.io.IOException;

/**
 * Samples for DataSources CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/DataSourcesCreate.json
     */
    /**
     * Sample code: DataSourcesCreate.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataSourcesCreate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager)
        throws IOException {
        manager.dataSources().define("AzTestDS774").withExistingWorkspace("OIAutoRest5123", "AzTest9724")
            .withProperties(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                "{\"LinkedResourceId\":\"/subscriptions/00000000-0000-0000-0000-00000000000/providers/microsoft.insights/eventtypes/management\"}",
                Object.class, SerializerEncoding.JSON))
            .withKind(DataSourceKind.AZURE_ACTIVITY_LOG).create();
    }
}
