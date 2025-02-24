
/**
 * Samples for DataProducts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProducts_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataProducts_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void
        dataProductsDeleteMaximumSetGen(com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataProducts().delete("aoiresourceGroupName", "dataproduct01", com.azure.core.util.Context.NONE);
    }
}
