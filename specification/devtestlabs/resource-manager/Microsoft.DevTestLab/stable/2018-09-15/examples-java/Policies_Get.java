/** Samples for Policies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_Get.json
     */
    /**
     * Sample code: Policies_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void policiesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .policies()
            .getWithResponse(
                "resourceGroupName",
                "{labName}",
                "{policySetName}",
                "{policyName}",
                null,
                com.azure.core.util.Context.NONE);
    }
}
