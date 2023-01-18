/** Samples for Labs ClaimAnyVm. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ClaimAnyVm.json
     */
    /**
     * Sample code: Labs_ClaimAnyVm.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsClaimAnyVm(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.labs().claimAnyVm("resourceGroupName", "{labName}", com.azure.core.util.Context.NONE);
    }
}
