/** Samples for Policies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_Delete.json
     */
    /**
     * Sample code: Policies_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void policiesDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .policies()
            .deleteWithResponse(
                "resourceGroupName", "{labName}", "{policySetName}", "{policyName}", com.azure.core.util.Context.NONE);
    }
}
