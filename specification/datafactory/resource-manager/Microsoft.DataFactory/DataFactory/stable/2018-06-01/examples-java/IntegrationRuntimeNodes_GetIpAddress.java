
/**
 * Samples for IntegrationRuntimeNodes GetIpAddress.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimeNodes_GetIpAddress.json
     */
    /**
     * Sample code: IntegrationRuntimeNodes_GetIpAddress.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimeNodesGetIpAddress(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimeNodes().getIpAddressWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", "Node_1", com.azure.core.util.Context.NONE);
    }
}
