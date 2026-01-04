
/**
 * Samples for NetworkSecurityPerimeterOperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspOperationStatusGet.json
     */
    /**
     * Sample code: NspOperationStatusGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspOperationStatusGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterOperationStatuses()
            .getWithResponse("location1", "operationId1", com.azure.core.util.Context.NONE);
    }
}
