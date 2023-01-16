/** Samples for Locations GetCapability. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Locations_GetCapability.json
     */
    /**
     * Sample code: Gets subscription-level properties and limits for Data Lake Store specified by resource location.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void getsSubscriptionLevelPropertiesAndLimitsForDataLakeStoreSpecifiedByResourceLocation(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.locations().getCapabilityWithResponse("EastUS2", com.azure.core.util.Context.NONE);
    }
}
