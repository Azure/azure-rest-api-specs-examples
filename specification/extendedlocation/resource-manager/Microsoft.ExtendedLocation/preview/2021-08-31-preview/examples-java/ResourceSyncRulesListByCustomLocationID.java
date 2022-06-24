import com.azure.core.util.Context;

/** Samples for ResourceSyncRules ListByCustomLocationId. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/ResourceSyncRulesListByCustomLocationID.json
     */
    /**
     * Sample code: List Resource Sync Rules by subscription.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void listResourceSyncRulesBySubscription(
        com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.resourceSyncRules().listByCustomLocationId("testresourcegroup", "customLocation01", Context.NONE);
    }
}
