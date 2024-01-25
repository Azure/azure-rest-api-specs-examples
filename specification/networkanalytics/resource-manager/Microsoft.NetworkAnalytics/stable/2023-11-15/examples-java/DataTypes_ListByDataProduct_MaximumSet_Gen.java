
/**
 * Samples for DataTypes ListByDataProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataTypes_ListByDataProduct_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataTypes_ListByDataProduct_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataTypesListByDataProductMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataTypes().listByDataProduct("aoiresourceGroupName", "dataproduct01",
            com.azure.core.util.Context.NONE);
    }
}
