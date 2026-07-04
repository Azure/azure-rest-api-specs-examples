
/**
 * Samples for NetworkSecurityPerimeterOperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspOperationStatusGet.json
     */
    /**
     * Sample code: NspOperationStatusGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspOperationStatusGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterOperationStatuses().getWithResponse("location1",
            "operationId1", com.azure.core.util.Context.NONE);
    }
}
