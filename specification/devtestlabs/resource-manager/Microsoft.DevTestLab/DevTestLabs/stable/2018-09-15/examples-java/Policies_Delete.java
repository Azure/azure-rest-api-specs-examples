
/**
 * Samples for Policies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Policies_Delete.json
     */
    /**
     * Sample code: Policies_Delete.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void policiesDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.policies().deleteWithResponse("resourceGroupName", "{labName}", "{policySetName}", "{policyName}",
            com.azure.core.util.Context.NONE);
    }
}
