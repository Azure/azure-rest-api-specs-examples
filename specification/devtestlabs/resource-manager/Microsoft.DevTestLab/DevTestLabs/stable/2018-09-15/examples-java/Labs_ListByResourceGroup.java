
/**
 * Samples for Labs ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Labs_ListByResourceGroup.json
     */
    /**
     * Sample code: Labs_ListByResourceGroup.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsListByResourceGroup(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.labs().listByResourceGroup("resourceGroupName", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
