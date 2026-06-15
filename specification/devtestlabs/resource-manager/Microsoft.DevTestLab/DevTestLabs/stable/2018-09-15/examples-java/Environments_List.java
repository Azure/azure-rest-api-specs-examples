
/**
 * Samples for Environments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Environments_List.json
     */
    /**
     * Sample code: Environments_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void environmentsList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.environments().list("resourceGroupName", "{labName}", "@me", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
