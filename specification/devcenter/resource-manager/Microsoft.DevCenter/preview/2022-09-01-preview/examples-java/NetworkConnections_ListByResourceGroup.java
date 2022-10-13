import com.azure.core.util.Context;

/** Samples for NetworkConnections ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/NetworkConnections_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkConnections_ListByResourceGroup.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsListByResourceGroup(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.networkConnections().listByResourceGroup("rg1", null, Context.NONE);
    }
}
