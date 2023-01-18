/** Samples for PrivateEndpointConnections ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateEndpointConnections_ListByAccount.json
     */
    /**
     * Sample code: PrivateEndpointConnections_ListByAccount.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void privateEndpointConnectionsListByAccount(
        com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .privateEndpointConnections()
            .listByAccount("SampleResourceGroup", "account1", null, com.azure.core.util.Context.NONE);
    }
}
