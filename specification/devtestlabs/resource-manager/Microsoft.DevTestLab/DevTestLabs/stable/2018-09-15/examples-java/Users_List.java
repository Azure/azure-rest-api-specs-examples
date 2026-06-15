
/**
 * Samples for Users List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Users_List.json
     */
    /**
     * Sample code: Users_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void usersList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.users().list("resourceGroupName", "{devtestlabName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
