
/**
 * Samples for DataTypes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataTypes_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataTypes_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void
        dataTypesDeleteMaximumSetGen(com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataTypes().delete("aoiresourceGroupName", "dataproduct01", "datatypename",
            com.azure.core.util.Context.NONE);
    }
}
