/** Samples for Policies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_List.json
     */
    /**
     * Sample code: Policies_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void policiesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .policies()
            .list(
                "resourceGroupName",
                "{labName}",
                "{policySetName}",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
