/** Samples for RemoteRenderingAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/remote-rendering/GetBySubscription.json
     */
    /**
     * Sample code: List remote rendering accounts by subscription.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void listRemoteRenderingAccountsBySubscription(
        com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager.remoteRenderingAccounts().list(com.azure.core.util.Context.NONE);
    }
}
