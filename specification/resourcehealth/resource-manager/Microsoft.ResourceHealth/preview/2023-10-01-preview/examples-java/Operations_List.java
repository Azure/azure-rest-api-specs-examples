
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/
     * Operations_List.json
     */
    /**
     * Sample code: GetOperationsList.
     * 
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void getOperationsList(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
