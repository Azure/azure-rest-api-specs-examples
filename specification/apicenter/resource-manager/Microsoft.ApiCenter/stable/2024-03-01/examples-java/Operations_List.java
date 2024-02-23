
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Operations_List.json
     */
    /**
     * Sample code: List Provider Operations.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void listProviderOperations(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
