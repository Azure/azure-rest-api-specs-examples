
/**
 * Samples for NetworkSecurityPerimeterConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/NetworkSecurityPerimeterConfigurationGet.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationGet.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        networkSecurityPerimeterConfigurationGet(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().getWithResponse("res4410", "sto8607",
            "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1", com.azure.core.util.Context.NONE);
    }
}
