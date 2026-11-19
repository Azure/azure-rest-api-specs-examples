
/**
 * Samples for NetworkGatewayResources Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/NetworkGatewayResources_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkGatewayResources_Delete_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        networkGatewayResourcesDeleteMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.networkGatewayResources().delete("rgconfluent", "myOrganization", "env-abc123", "gw-def456",
            com.azure.core.util.Context.NONE);
    }
}
