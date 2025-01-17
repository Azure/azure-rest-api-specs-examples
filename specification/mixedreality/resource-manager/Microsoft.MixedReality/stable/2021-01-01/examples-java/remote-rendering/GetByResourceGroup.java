
/**
 * Samples for RemoteRenderingAccounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/remote-rendering/
     * GetByResourceGroup.json
     */
    /**
     * Sample code: List remote rendering accounts by resource group.
     * 
     * @param manager Entry point to MixedRealityManager.
     */
    public static void
        listRemoteRenderingAccountsByResourceGroup(com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager.remoteRenderingAccounts().listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
