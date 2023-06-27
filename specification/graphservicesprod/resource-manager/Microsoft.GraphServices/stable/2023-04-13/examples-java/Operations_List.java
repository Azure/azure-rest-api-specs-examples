/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Operations_List.json
     */
    /**
     * Sample code: Get list of operations.
     *
     * @param manager Entry point to GraphServicesManager.
     */
    public static void getListOfOperations(com.azure.resourcemanager.graphservices.GraphServicesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
