
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void operationsList(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
