
/**
 * Samples for NetworkSecurityPerimeterConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * NetworkSecurityPerimeterConfigurationList.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkSecurityPerimeterConfigurationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getNetworkSecurityPerimeterConfigurations().list("res4410",
            "sto8607", com.azure.core.util.Context.NONE);
    }
}
