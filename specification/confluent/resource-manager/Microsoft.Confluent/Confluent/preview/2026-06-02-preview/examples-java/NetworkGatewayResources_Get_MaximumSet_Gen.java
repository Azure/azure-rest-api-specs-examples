
/**
 * Samples for NetworkGatewayResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/NetworkGatewayResources_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkGatewayResources_Get_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        networkGatewayResourcesGetMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.networkGatewayResources().getWithResponse("rgconfluent", "myOrganization", "env-abc123", "gw-def456",
            com.azure.core.util.Context.NONE);
    }
}
