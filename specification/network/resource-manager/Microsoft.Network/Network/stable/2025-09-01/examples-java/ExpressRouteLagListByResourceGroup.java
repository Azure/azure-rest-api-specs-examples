
/**
 * Samples for ExpressRouteLags ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteLagListByResourceGroup.json
     */
    /**
     * Sample code: List express route lags by resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteLagsByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
