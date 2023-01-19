/** Samples for RemoteRenderingAccounts ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/remote-rendering/ListKeys.json
     */
    /**
     * Sample code: List remote rendering account key.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void listRemoteRenderingAccountKey(
        com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager
            .remoteRenderingAccounts()
            .listKeysWithResponse("MyResourceGroup", "MyAccount", com.azure.core.util.Context.NONE);
    }
}
