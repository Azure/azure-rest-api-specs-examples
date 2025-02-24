
import com.azure.resourcemanager.networkanalytics.models.DataTypeProperties;
import com.azure.resourcemanager.networkanalytics.models.DataTypeState;

/**
 * Samples for DataTypes Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataTypes_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataTypes_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void
        dataTypesCreateMaximumSetGen(com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataTypes().define("datatypename").withExistingDataProduct("aoiresourceGroupName", "dataproduct01")
            .withProperties(new DataTypeProperties().withState(DataTypeState.fromString("STARTED"))
                .withStorageOutputRetention(27).withDatabaseCacheRetention(23).withDatabaseRetention(6))
            .create();
    }
}
