/** Samples for KustoPools ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListByWorkspace.json
     */
    /**
     * Sample code: List Kusto pools in a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listKustoPoolsInAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .listByWorkspaceWithResponse("kustorptest", "kustorptest", com.azure.core.util.Context.NONE);
    }
}
