
/**
 * Samples for Labs ListVhds.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Labs_ListVhds.json
     */
    /**
     * Sample code: Labs_ListVhds.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsListVhds(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.labs().listVhds("resourceGroupName", "{labName}", com.azure.core.util.Context.NONE);
    }
}
