
/**
 * Samples for Labs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Labs_Delete.json
     */
    /**
     * Sample code: Labs_Delete.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.labs().delete("resourceGroupName", "{labName}", com.azure.core.util.Context.NONE);
    }
}
