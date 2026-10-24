
/**
 * Samples for ParameterGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ParameterGroups_Get.json
     */
    /**
     * Sample code: Get a HorizonDB parameter group.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBParameterGroup(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.parameterGroups().getByResourceGroupWithResponse("example-resource-group", "production-parameters",
            com.azure.core.util.Context.NONE);
    }
}
