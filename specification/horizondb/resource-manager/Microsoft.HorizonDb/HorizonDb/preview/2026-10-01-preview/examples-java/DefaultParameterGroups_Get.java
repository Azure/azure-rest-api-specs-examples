
/**
 * Samples for DefaultParameterGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/DefaultParameterGroups_Get.json
     */
    /**
     * Sample code: Get a default HorizonDB parameter group by subscription, location, and name.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getADefaultHorizonDBParameterGroupBySubscriptionLocationAndName(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.defaultParameterGroups().getWithResponse("eastus2", "default-postgresql-17-provisioned",
            com.azure.core.util.Context.NONE);
    }
}
