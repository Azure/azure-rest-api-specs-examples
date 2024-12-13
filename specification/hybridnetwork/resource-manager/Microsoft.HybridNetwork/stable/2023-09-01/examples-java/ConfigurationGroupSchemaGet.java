
/**
 * Samples for ConfigurationGroupSchemas Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ConfigurationGroupSchemaGet.json
     */
    /**
     * Sample code: Get a networkFunctionDefinition group resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getANetworkFunctionDefinitionGroupResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.configurationGroupSchemas().getWithResponse("rg1", "testPublisher", "testConfigurationGroupSchema",
            com.azure.core.util.Context.NONE);
    }
}
