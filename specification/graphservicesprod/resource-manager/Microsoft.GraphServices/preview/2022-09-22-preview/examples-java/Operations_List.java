/** Samples for Operation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/preview/2022-09-22-preview/examples/Operations_List.json
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
