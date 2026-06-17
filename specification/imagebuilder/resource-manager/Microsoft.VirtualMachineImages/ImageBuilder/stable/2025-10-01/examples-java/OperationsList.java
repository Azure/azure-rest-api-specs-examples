
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/OperationsList.json
     */
    /**
     * Sample code: Retrieve operations list.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void retrieveOperationsList(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
