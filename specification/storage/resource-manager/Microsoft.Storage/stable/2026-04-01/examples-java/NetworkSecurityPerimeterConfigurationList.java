
/**
 * Samples for NetworkSecurityPerimeterConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/NetworkSecurityPerimeterConfigurationList.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        networkSecurityPerimeterConfigurationList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().list("res4410", "sto8607",
            com.azure.core.util.Context.NONE);
    }
}
