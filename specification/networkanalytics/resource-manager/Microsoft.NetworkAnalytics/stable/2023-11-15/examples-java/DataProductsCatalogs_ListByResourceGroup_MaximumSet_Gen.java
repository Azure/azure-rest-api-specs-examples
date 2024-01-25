
/**
 * Samples for DataProductsCatalogs ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProductsCatalogs_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataProductsCatalogs_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataProductsCatalogsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataProductsCatalogs().listByResourceGroup("aoiresourceGroupName", com.azure.core.util.Context.NONE);
    }
}
