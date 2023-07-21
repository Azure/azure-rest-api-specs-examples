/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetOperations.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void operationsList(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
