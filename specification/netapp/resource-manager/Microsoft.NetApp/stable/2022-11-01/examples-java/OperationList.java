/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/OperationList.json
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
