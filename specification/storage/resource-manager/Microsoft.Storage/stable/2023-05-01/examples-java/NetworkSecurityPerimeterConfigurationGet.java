
/**
 * Samples for NetworkSecurityPerimeterConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * NetworkSecurityPerimeterConfigurationGet.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkSecurityPerimeterConfigurationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getNetworkSecurityPerimeterConfigurations().getWithResponse(
            "res4410", "sto8607", "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1",
            com.azure.core.util.Context.NONE);
    }
}
