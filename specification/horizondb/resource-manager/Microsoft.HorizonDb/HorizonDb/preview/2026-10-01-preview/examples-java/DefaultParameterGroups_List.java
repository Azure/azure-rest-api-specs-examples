
/**
 * Samples for DefaultParameterGroups ListBySubscriptionAndRegion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/DefaultParameterGroups_List.json
     */
    /**
     * Sample code: List default HorizonDB parameter groups by subscription and region.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listDefaultHorizonDBParameterGroupsBySubscriptionAndRegion(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.defaultParameterGroups().listBySubscriptionAndRegion("eastus2", com.azure.core.util.Context.NONE);
    }
}
