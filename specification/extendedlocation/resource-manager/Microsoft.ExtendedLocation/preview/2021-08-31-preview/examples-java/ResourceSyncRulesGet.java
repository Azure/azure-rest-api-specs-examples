import com.azure.core.util.Context;

/** Samples for ResourceSyncRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/ResourceSyncRulesGet.json
     */
    /**
     * Sample code: Get Custom Location.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void getCustomLocation(com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager
            .resourceSyncRules()
            .getWithResponse("testresourcegroup", "customLocation01", "resourceSyncRule01", Context.NONE);
    }
}
