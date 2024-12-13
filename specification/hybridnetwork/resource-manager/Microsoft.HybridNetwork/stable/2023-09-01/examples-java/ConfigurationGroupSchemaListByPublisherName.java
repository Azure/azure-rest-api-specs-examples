
/**
 * Samples for ConfigurationGroupSchemas ListByPublisher.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ConfigurationGroupSchemaListByPublisherName.json
     */
    /**
     * Sample code: Get networkFunctionDefinition groups under publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getNetworkFunctionDefinitionGroupsUnderPublisherResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.configurationGroupSchemas().listByPublisher("rg1", "testPublisher", com.azure.core.util.Context.NONE);
    }
}
