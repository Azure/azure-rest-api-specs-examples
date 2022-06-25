import com.azure.core.util.Context;

/** Samples for ResourceSyncRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/ResourceSyncRulesDelete.json
     */
    /**
     * Sample code: Delete Resource Sync Rule.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void deleteResourceSyncRule(
        com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager
            .resourceSyncRules()
            .deleteWithResponse("testresourcegroup", "customLocation01", "resourceSyncRule01", Context.NONE);
    }
}
