
/**
 * Samples for Connections ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/listConnectionsByResourceGroup.json
     */
    /**
     * Sample code: Gets connections in a resource group.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        getsConnectionsInAResourceGroup(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.connections().listByResourceGroup("testRG", com.azure.core.util.Context.NONE);
    }
}
