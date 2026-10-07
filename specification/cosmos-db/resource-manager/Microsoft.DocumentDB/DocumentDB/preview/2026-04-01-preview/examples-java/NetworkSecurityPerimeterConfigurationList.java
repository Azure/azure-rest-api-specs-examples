
/**
 * Samples for NetworkSecurityPerimeterConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/NetworkSecurityPerimeterConfigurationList.json
     */
    /**
     * Sample code: NamspaceNetworkSecurityPerimeterConfigurationList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        namspaceNetworkSecurityPerimeterConfigurationList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().list("res4410", "cosmosTest",
            com.azure.core.util.Context.NONE);
    }
}
