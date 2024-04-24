
/**
 * Samples for NetworkConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/NetworkConnections_Delete
     * .json
     */
    /**
     * Sample code: NetworkConnections_Delete.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.networkConnections().delete("rg1", "eastusnetwork", com.azure.core.util.Context.NONE);
    }
}
