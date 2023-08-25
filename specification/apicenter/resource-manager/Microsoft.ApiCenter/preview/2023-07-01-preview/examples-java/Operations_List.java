/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/preview/2023-07-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to ApiCenterManager.
     */
    public static void operationsList(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
