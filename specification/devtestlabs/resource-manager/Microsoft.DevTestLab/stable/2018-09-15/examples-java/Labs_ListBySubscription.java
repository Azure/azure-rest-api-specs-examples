/** Samples for Labs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ListBySubscription.json
     */
    /**
     * Sample code: Labs_ListBySubscription.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsListBySubscription(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.labs().list(null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
