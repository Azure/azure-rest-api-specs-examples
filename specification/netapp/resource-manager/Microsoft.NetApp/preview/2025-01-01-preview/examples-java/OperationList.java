
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/OperationList.json
     */
    /**
     * Sample code: OperationList.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void operationList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
