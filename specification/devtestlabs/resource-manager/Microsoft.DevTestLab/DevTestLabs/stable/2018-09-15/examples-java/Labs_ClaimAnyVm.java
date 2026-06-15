
/**
 * Samples for Labs ClaimAnyVm.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Labs_ClaimAnyVm.json
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
