
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for DataProducts ListRolesAssignments.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProducts_ListRolesAssignments_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataProducts_ListRolesAssignments_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataProductsListRolesAssignmentsMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) throws IOException {
        manager.dataProducts()
            .listRolesAssignmentsWithResponse("aoiresourceGroupName", "dataproduct01", SerializerFactory
                .createDefaultManagementSerializerAdapter().deserialize("{}", Object.class, SerializerEncoding.JSON),
                com.azure.core.util.Context.NONE);
    }
}
