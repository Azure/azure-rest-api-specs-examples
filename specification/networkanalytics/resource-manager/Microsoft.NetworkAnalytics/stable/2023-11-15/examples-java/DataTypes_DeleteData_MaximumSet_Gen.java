
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for DataTypes DeleteData.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataTypes_DeleteData_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataTypes_DeleteData_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataTypesDeleteDataMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) throws IOException {
        manager.dataTypes()
            .deleteData("aoiresourceGroupName", "dataproduct01", "datatypename", SerializerFactory
                .createDefaultManagementSerializerAdapter().deserialize("{}", Object.class, SerializerEncoding.JSON),
                com.azure.core.util.Context.NONE);
    }
}
