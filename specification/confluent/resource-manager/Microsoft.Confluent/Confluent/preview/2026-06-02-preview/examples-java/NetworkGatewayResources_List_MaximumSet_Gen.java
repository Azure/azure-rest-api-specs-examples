
/**
 * Samples for NetworkGatewayResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/NetworkGatewayResources_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkGatewayResources_List_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        networkGatewayResourcesListMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.networkGatewayResources().list("rgconfluent", "myOrganization", "env-abc123",
            com.azure.core.util.Context.NONE);
    }
}
